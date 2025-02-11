package utils

import (
	"crypto/ed25519"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"strings"

	"github.com/mr-tron/base58"
	"golang.org/x/crypto/curve25519"
)

// By default, we use "did:peer:2". The user can override by passing --method=did:scid:pr:2 (or any other prefix).
var methodPrefix string

func init() {
	// Provide a command-line flag for customizing the DID method prefix.
	flag.StringVar(&methodPrefix, "method", "did:peer:2", "Method prefix, e.g. did:peer:2 or did:scid:pr:2")
}

// ServiceProfile holds additional profile metadata for the service endpoint.
type ServiceProfile struct {
	Profile   string `json:"profile"`
	URI       string `json:"uri"`
	Integrity string `json:"integrity"`
}

// Service represents an entry in the "services" array for a DID peer config.
type Service struct {
	ID              string         `json:"id"`
	Type            string         `json:"type"`
	ServiceEndpoint ServiceProfile `json:"serviceEndpoint"`
}

// Peer2ConfigFile is just an example of how you might wrap the array of services.
type Peer2ConfigFile struct {
	Services []Service `json:"services"`
}

// generateDidPeer2 creates a did:peer:2-style (or did:scid:pr:2, etc.) DID from 2 generated keys (Ed25519, X25519) + the provided services.
func GenerateDidPeer2(cfg Peer2ConfigFile) (string, error) {
	// Generate Ed25519
	edPub, edPriv, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		return "", fmt.Errorf("generate ed25519: %w", err)
	}
	edPubMB := "z" + base58.Encode(edPub)

	// Generate X25519
	var x25519Priv [32]byte
	if _, err := rand.Read(x25519Priv[:]); err != nil {
		return "", fmt.Errorf("rand x25519 priv: %w", err)
	}
	x25519Pub, err := curve25519.X25519(x25519Priv[:], curve25519.Basepoint)
	if err != nil {
		return "", fmt.Errorf("curve25519 derive pub: %w", err)
	}
	x25519PubMB := "z" + base58.Encode(x25519Pub)

	// Print the keys
	fmt.Println("=== Generated Keys ===")
	fmt.Println("Ed25519 Private Key (hex):", hex.EncodeToString(edPriv))
	fmt.Println("Ed25519 Public Key  (MB): ", edPubMB)
	fmt.Println("X25519 Private Key  (hex):", hex.EncodeToString(x25519Priv[:]))
	fmt.Println("X25519 Public Key   (MB): ", x25519PubMB)
	fmt.Println()

	// Build DID e.g.: did:peer:2.Vz6Mkj3PU...Ez6LSg8zQ...
	var sb strings.Builder
	sb.WriteString(methodPrefix) // e.g. "did:peer:2" or "did:scid:pr:2"
	sb.WriteString(".V")
	sb.WriteString(edPubMB)
	sb.WriteString(".E")
	sb.WriteString(x25519PubMB)

	// For each service => .S + base64url( abbreviateService(...) )
	for _, svc := range cfg.Services {
		abbrev := abbreviateService(svc)
		svcBytes, err := json.Marshal(abbrev)
		if err != nil {
			return "", fmt.Errorf("marshal service: %w", err)
		}
		b64 := base64.RawURLEncoding.EncodeToString(svcBytes)
		sb.WriteString(".S")
		sb.WriteString(b64)
	}

	return sb.String(), nil
}

// abbreviateService tries to shorten fields like "type" -> "t", "serviceEndpoint"->"s", etc.
var fullToAbbreviation = map[string]string{
	"type":             "t",
	"serviceEndpoint":  "s",
	"routingKeys":      "r",
	"accept":           "a",
	"DIDCommMessaging": "dm", // example for DIDComm v2
}

func abbreviateService(obj interface{}) interface{} {
	switch val := obj.(type) {
	case map[string]interface{}:
		newMap := map[string]interface{}{}
		for k, v := range val {
			abbrKey := fullToAbbreviation[k]
			if abbrKey == "" {
				abbrKey = k // no abbreviation
			}
			if k == "type" {
				// Possibly abbreviate the value if it's "DIDCommMessaging" => "dm"
				if strval, ok := v.(string); ok {
					if code, found := fullToAbbreviation[strval]; found {
						v = code
					}
				}
			}
			newMap[abbrKey] = abbreviateService(v)
		}
		return newMap
	case []interface{}:
		newArr := make([]interface{}, len(val))
		for i, elem := range val {
			newArr[i] = abbreviateService(elem)
		}
		return newArr
	default:
		return val
	}
}

// resolveDidPeer2 takes a DID (default "did:peer:2", or maybe "did:scid:pr:2") and decodes it into a DID Document
func resolveDidPeer2(didStr string) (map[string]interface{}, error) {
	// Check the prefix
	if !strings.HasPrefix(didStr, methodPrefix) {
		return nil, fmt.Errorf("DID %q does not match method prefix %q", didStr, methodPrefix)
	}

	doc := map[string]interface{}{
		"@context": []interface{}{
			"https://www.w3.org/ns/did/v1",
			"https://w3id.org/security/multikey/v1",
		},
		"id": didStr,
	}

	vmList := []map[string]interface{}{}
	auth := []string{}
	assertion := []string{}
	keyAgreement := []string{}
	capInv := []string{}
	capDel := []string{}

	var services []map[string]interface{}
	serviceCount := 0

	// remove prefix
	body := strings.TrimPrefix(didStr, methodPrefix)
	parts := strings.Split(body, ".")

	keyIndex := 1

	for i, seg := range parts {
		if i == 0 {
			// usually empty
			continue
		}
		if len(seg) < 2 {
			return nil, fmt.Errorf("malformed segment '%s'", seg)
		}
		purpose := seg[0] // e.g. 'V','E','S' etc.
		rest := seg[1:]

		switch purpose {
		case 'V', 'A', 'E', 'I', 'D':
			keyID := fmt.Sprintf("#key-%d", keyIndex)
			keyIndex++
			vm := map[string]interface{}{
				"id":                 keyID,
				"type":               "Multikey",
				"controller":         didStr,
				"publicKeyMultibase": rest,
			}
			vmList = append(vmList, vm)
			switch purpose {
			case 'V':
				auth = append(auth, keyID)
			case 'A':
				assertion = append(assertion, keyID)
			case 'E':
				keyAgreement = append(keyAgreement, keyID)
			case 'I':
				capInv = append(capInv, keyID)
			case 'D':
				capDel = append(capDel, keyID)
			}

		case 'S':
			decoded, err := base64.RawURLEncoding.DecodeString(rest)
			if err != nil {
				return nil, fmt.Errorf("decode base64 service: %w", err)
			}
			var raw interface{}
			if err := json.Unmarshal(decoded, &raw); err != nil {
				return nil, fmt.Errorf("unmarshal service: %w", err)
			}
			exp := expandService(raw)
			svcObj, ok := exp.(map[string]interface{})
			if !ok {
				return nil, errors.New("decoded service is not an object")
			}
			// if no "id", auto-assign
			if _, found := svcObj["id"]; !found {
				if serviceCount == 0 {
					svcObj["id"] = "#service"
				} else {
					svcObj["id"] = fmt.Sprintf("#service-%d", serviceCount)
				}
			}
			serviceCount++
			services = append(services, svcObj)

		default:
			return nil, fmt.Errorf("unknown purpose code '%c' in segment '%s'", purpose, seg)
		}
	}

	if len(vmList) > 0 {
		doc["verificationMethod"] = vmList
	}
	if len(auth) > 0 {
		doc["authentication"] = auth
	}
	if len(assertion) > 0 {
		doc["assertionMethod"] = assertion
	}
	if len(keyAgreement) > 0 {
		doc["keyAgreement"] = keyAgreement
	}
	if len(capInv) > 0 {
		doc["capabilityInvocation"] = capInv
	}
	if len(capDel) > 0 {
		doc["capabilityDelegation"] = capDel
	}
	if len(services) > 0 {
		doc["service"] = services
	}

	return doc, nil
}

var abbreviationToFull = map[string]string{
	"t":  "type",
	"s":  "serviceEndpoint",
	"r":  "routingKeys",
	"a":  "accept",
	"dm": "DIDCommMessaging",
}

func expandService(obj interface{}) interface{} {
	switch val := obj.(type) {
	case map[string]interface{}:
		newMap := map[string]interface{}{}
		for k, v := range val {
			expandedKey := k
			if full, ok := abbreviationToFull[k]; ok {
				expandedKey = full
			}
			if expandedKey == "type" {
				if strval, ok := v.(string); ok {
					if strval == "dm" {
						v = "DIDCommMessaging"
					}
				}
			}
			newMap[expandedKey] = expandService(v)
		}
		return newMap
	case []interface{}:
		newArr := make([]interface{}, len(val))
		for i, elem := range val {
			newArr[i] = expandService(elem)
		}
		return newArr
	default:
		return val
	}
}

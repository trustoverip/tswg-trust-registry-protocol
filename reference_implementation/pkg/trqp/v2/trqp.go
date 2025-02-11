package v2

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v4"

	"github.com/andorsk/tswg-trust-registry-protocol/reference_implementation/gen/trqp/v2"
	"github.com/andorsk/tswg-trust-registry-protocol/reference_implementation/pkg/utils"
)

var ecdsaPrivateKey *ecdsa.PrivateKey

func init() {
	// Generate an ephemeral ECDSA P-256 private key on startup.
	// In production, you’d likely load this key from a secure store or file.
	var err error
	ecdsaPrivateKey, err = ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		panic("failed to generate ECDSA private key: " + err.Error())
	}
}

type TRQPHandler struct {
	Registry *utils.TrustRegistry
}

func (impl *TRQPHandler) CheckEcosystemRecognition(
	w http.ResponseWriter,
	r *http.Request,
	egfDid string,
	ecosystemId string,
	params trqp.CheckEcosystemRecognitionParams,
) {
	ecosystem, err := impl.Registry.GetEcosystemByDID(egfDid)
	if err != nil {
		writeError(w, http.StatusNotFound, "Ecosystem not found", err.Error())
		return
	}

	if ecosystem.Metadata.DID != egfDid {
		writeError(w, http.StatusNotFound, "Ecosystem ID mismatch",
			"The ecosystem DID does not match the provided ecosystem_id.")
		return
	}

	var isRecognized = false

	for _, entry := range ecosystem.RecognitionEntries {
		if entry.DID == ecosystemId && entry.Status.Active {
			isRecognized = true
			break

		}
	}

	// Here we assume that an ecosystem is recognized if its metadata status is "active"

	recognitionPayload := map[string]interface{}{
		"ecosystem_did": ecosystem.Metadata.DID,
		"egf_did":       egfDid,
		"recognized":    isRecognized,
		"timestamp":     time.Now().UTC().Format(time.RFC3339),
		"nonce":         params.Nonce,
	}
	payloadBytes, _ := json.Marshal(recognitionPayload)

	jws, err := signPayload(payloadBytes)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "Failed to sign payload", err.Error())
		return
	}

	resp := trqp.RecognitionResponseJWS{Jws: jws}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func (impl *TRQPHandler) CheckAuthorizationStatus(
	w http.ResponseWriter,
	r *http.Request,
	entityId string,
	authorizationId string,
	egfDid string,
	params trqp.CheckAuthorizationStatusParams,
) {
	ecosystem, err := impl.Registry.GetEcosystemByDID(egfDid)
	if err != nil {
		writeError(w, http.StatusNotFound, "Ecosystem not found", err.Error())
		return
	}

	var found *utils.AuthorizationEntry
	for i, authEntry := range ecosystem.AuthorizationEntries {
		fmt.Println(authEntry.Authorization, authorizationId)
		if authEntry.ID == authorizationId && authEntry.DID == entityId {
			found = &ecosystem.AuthorizationEntries[i]
			break
		}
	}

	if found == nil {
		writeError(w, http.StatusNotFound, "Authorization not found",
			"No matching authorization entry for entity.")
		return
	}

	// If no status was provided in the authorization entry, assume it is active.
	authorized := true
	if found.Status.Active != true {
		authorized = (found.Status.Active == true)
	}

	authorizationPayload := map[string]interface{}{
		"entity_id":        entityId,
		"authorization_id": authorizationId,
		"authorized":       authorized,
		"timestamp":        time.Now().UTC().Format(time.RFC3339),
		"nonce":            params.Nonce,
	}
	payloadBytes, _ := json.Marshal(authorizationPayload)

	jws, err := signPayload(payloadBytes)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "Failed to sign payload", err.Error())
		return
	}

	resp := trqp.AuthorizationResponseJWS{Jws: jws}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func (impl *TRQPHandler) GetTrustRegistryMetadata(w http.ResponseWriter, r *http.Request) {
	metadataBytes, err := json.Marshal(impl.Registry.Metadata)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "Metadata serialization error", err.Error())
		return
	}

	jws, err := signPayload(metadataBytes)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "Failed to sign payload", err.Error())
		return
	}

	resp := trqp.TrustRegistryMetadataJWS{Jws: jws}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

// signPayload creates a compact JWS string (header.payload.signature)
// using ES256 (ECDSA with P-256 and SHA-256).
func signPayload(payload []byte) (string, error) {
	// Manually build a JWT-like structure.
	header := base64.RawURLEncoding.EncodeToString([]byte(`{"alg":"ES256","typ":"JWT"}`))
	payloadEncoded := base64.RawURLEncoding.EncodeToString(payload)
	unsigned := header + "." + payloadEncoded

	signMethod := jwt.GetSigningMethod("ES256")
	signature, err := signMethod.Sign(unsigned, ecdsaPrivateKey)
	if err != nil {
		return "", err
	}
	return unsigned + "." + signature, nil
}

func writeError(w http.ResponseWriter, code int, errShort, details string) {
	w.WriteHeader(code)
	resp := trqp.Error{
		Error:   errShort,
		Details: &details,
	}
	json.NewEncoder(w).Encode(resp)
}

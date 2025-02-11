package utils

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Test DID Generation
func TestGenerateDidPeer2_Success(t *testing.T) {
	cfg := Peer2ConfigFile{
		Services: []Service{
			{
				ID:   "service1",
				Type: "DIDCommMessaging",
				ServiceEndpoint: ServiceProfile{
					Profile:   "test-profile",
					URI:       "https://example.com",
					Integrity: "sha256:xyz",
				},
			},
		},
	}

	did, err := GenerateDidPeer2(cfg)
	assert.NoError(t, err)
	assert.NotEmpty(t, did)

	// DID should start with methodPrefix
	assert.True(t, strings.HasPrefix(did, methodPrefix), "DID should start with methodPrefix")

	// Check if the DID contains expected sections
	assert.Contains(t, did, ".V", "Should contain verification key")
	assert.Contains(t, did, ".E", "Should contain encryption key")
	assert.Contains(t, ".S", did, "Should contain service section")
}

// Test DID Resolution Success
func TestResolveDidPeer2_Success(t *testing.T) {
	did := "did:peer:2.Vz6Mkj3PUBg2waEz6LSg8zQ.Ez6LSg8zQxSSz4yz.SeyJ0IjoiRElEQ29tbU1lc3NhZ2luZyIsInMiOiJodHRwczovL2V4YW1wbGUuY29tIn0"

	doc, err := resolveDidPeer2(did)
	assert.NoError(t, err)
	assert.NotNil(t, doc)

	// Verify DID Document structure
	assert.Equal(t, did, doc["id"])
	assert.Contains(t, doc, "verificationMethod")
	assert.Contains(t, doc, "service")

	// Verify service resolution
	services, ok := doc["service"].([]map[string]interface{})
	assert.True(t, ok)
	assert.Len(t, services, 1)
	assert.Equal(t, "DIDCommMessaging", services[0]["type"])
	assert.Equal(t, "https://example.com", services[0]["serviceEndpoint"])
}

// Test Resolving a DID with a Different Prefix
func TestResolveDidPeer2_WrongPrefix(t *testing.T) {
	did := "did:wrong:2.Vz6Mkj3PUBg2waEz6LSg8zQ"

	_, err := resolveDidPeer2(did)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "does not match method prefix")
}

// Test Resolving an Invalid DID
func TestResolveDidPeer2_Malformed(t *testing.T) {
	did := "did:peer:2.Vz6Mkj3PUBg2waEz6LSg8zQ.S"

	_, err := resolveDidPeer2(did)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "malformed segment")
}

// Test Service Expansion
func TestExpandService(t *testing.T) {
	compactService := map[string]interface{}{
		"t": "dm",
		"s": "https://example.com",
	}

	expanded := expandService(compactService).(map[string]interface{})
	assert.Equal(t, "DIDCommMessaging", expanded["type"])
	assert.Equal(t, "https://example.com", expanded["serviceEndpoint"])
}

// Test Service Abbreviation
func TestAbbreviateService(t *testing.T) {
	fullService := map[string]interface{}{
		"type":            "DIDCommMessaging",
		"serviceEndpoint": "https://example.com",
	}

	abbreviated := abbreviateService(fullService).(map[string]interface{})
	assert.Equal(t, "dm", abbreviated["t"])
	assert.Equal(t, "https://example.com", abbreviated["s"])
}

// Test Abbreviate & Expand Roundtrip
func TestAbbreviateAndExpandService_Roundtrip(t *testing.T) {
	original := map[string]interface{}{
		"type":            "DIDCommMessaging",
		"serviceEndpoint": "https://example.com",
	}

	abbreviated := abbreviateService(original).(map[string]interface{})
	expanded := expandService(abbreviated).(map[string]interface{})

	assert.Equal(t, original, expanded, "Expanded service should match original")
}

// Test Marshaling Abbreviated Service
func TestAbbreviateService_Marshal(t *testing.T) {
	fullService := Service{
		ID:   "service1",
		Type: "DIDCommMessaging",
		ServiceEndpoint: ServiceProfile{
			Profile:   "test-profile",
			URI:       "https://example.com",
			Integrity: "sha256:xyz",
		},
	}

	abbreviated := abbreviateService(fullService)
	data, err := json.Marshal(abbreviated)
	assert.NoError(t, err)
	assert.Contains(t, string(data), `"t":"dm"`)
	assert.Contains(t, string(data), `"s":"https://example.com"`)
}

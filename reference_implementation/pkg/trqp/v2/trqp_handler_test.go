package v2

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/andorsk/tswg-trust-registry-protocol/reference_implementation/gen/trqp/v2"
	"github.com/andorsk/tswg-trust-registry-protocol/reference_implementation/pkg/utils"
	"github.com/stretchr/testify/assert"
)

func generateTestPrivateKey() *ecdsa.PrivateKey {
	privKey, _ := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	return privKey
}

// Mock Trust Registry
func setupMockRegistry() *utils.TrustRegistry {
	return &utils.TrustRegistry{
		Ecosystems: []utils.Ecosystem{
			{
				Metadata: utils.EcosystemMetadata{
					DID:    "did:example:123",
					Status: utils.Status{Active: true},
				},
				AuthorizationEntries: []utils.AuthorizationEntry{
					{
						Authorization: "auth-1",
						Status:        utils.Status{Active: true},
					},
				},
			},
		},
	}
}

func TestCheckEcosystemRecognition_Success(t *testing.T) {
	handler := &TRQPHandler{Registry: setupMockRegistry()}

	req := httptest.NewRequest("GET", "/ecosystem-recognition", nil)
	w := httptest.NewRecorder()

	var nonce = "random-nonce"

	params := trqp.CheckEcosystemRecognitionParams{
		Nonce: &nonce,
	}

	handler.CheckEcosystemRecognition(w, req, "did:example:123", "did:example:123", params)

	resp := w.Result()
	defer resp.Body.Close()

	assert.Equal(t, http.StatusOK, resp.StatusCode)

	var responseJWS trqp.RecognitionResponseJWS
	err := json.NewDecoder(resp.Body).Decode(&responseJWS)
	assert.NoError(t, err)
	assert.NotEmpty(t, responseJWS.Jws)
}

func TestCheckEcosystemRecognition_NotFound(t *testing.T) {
	handler := &TRQPHandler{Registry: setupMockRegistry()}

	req := httptest.NewRequest("GET", "/ecosystem-recognition", nil)
	w := httptest.NewRecorder()

	var nonce = "random-nonce"
	params := trqp.CheckEcosystemRecognitionParams{
		Nonce: &nonce,
	}

	handler.CheckEcosystemRecognition(w, req, "did:example:999", "did:example:999", params)

	resp := w.Result()
	defer resp.Body.Close()

	assert.Equal(t, http.StatusNotFound, resp.StatusCode)
}

func TestCheckAuthorizationStatus_Success(t *testing.T) {
	handler := &TRQPHandler{Registry: setupMockRegistry()}

	req := httptest.NewRequest("GET", "/authorization-status", nil)
	w := httptest.NewRecorder()

	params := trqp.CheckAuthorizationStatusParams{
		Nonce: "random-nonce",
	}

	handler.CheckAuthorizationStatus(w, req, "entity-1", "auth-1", "did:example:123", params)

	resp := w.Result()
	defer resp.Body.Close()

	assert.Equal(t, http.StatusOK, resp.StatusCode)

	var responseJWS trqp.AuthorizationResponseJWS
	err := json.NewDecoder(resp.Body).Decode(&responseJWS)
	assert.NoError(t, err)
	assert.NotEmpty(t, responseJWS.Jws)
}

func TestCheckAuthorizationStatus_NotFound(t *testing.T) {
	handler := &TRQPHandler{Registry: setupMockRegistry()}

	req := httptest.NewRequest("GET", "/authorization-status", nil)
	w := httptest.NewRecorder()

	params := trqp.CheckAuthorizationStatusParams{
		Nonce: "random-nonce",
	}

	handler.CheckAuthorizationStatus(w, req, "entity-1", "auth-999", "did:example:123", params)

	resp := w.Result()
	defer resp.Body.Close()

	assert.Equal(t, http.StatusNotFound, resp.StatusCode)
}

func TestGetTrustRegistryMetadata_Success(t *testing.T) {
	handler := &TRQPHandler{Registry: setupMockRegistry()}

	req := httptest.NewRequest("GET", "/trust-registry-metadata", nil)
	w := httptest.NewRecorder()

	handler.GetTrustRegistryMetadata(w, req)

	resp := w.Result()
	defer resp.Body.Close()

	assert.Equal(t, http.StatusOK, resp.StatusCode)

	var responseJWS trqp.TrustRegistryMetadataJWS
	err := json.NewDecoder(resp.Body).Decode(&responseJWS)
	assert.NoError(t, err)
	assert.NotEmpty(t, responseJWS.Jws)
}

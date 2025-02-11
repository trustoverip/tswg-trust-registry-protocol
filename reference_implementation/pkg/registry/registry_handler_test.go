package registry

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/andorsk/tswg-trust-registry-protocol/reference_implementation/gen/admin"
	"github.com/andorsk/tswg-trust-registry-protocol/reference_implementation/pkg/utils"
	"github.com/stretchr/testify/assert"
)

// --- Helper functions for pointer values ---
func StringPointer(s string) *string { return &s }
func BoolPointer(b bool) *bool       { return &b }

// ----------------- Handler Tests -----------------

func TestHandlerCreateEcosystem_Success(t *testing.T) {
	svc := newTestTrustRegistryService()
	handler := NewTrustRegistryHandlers(svc)

	eco := utils.Ecosystem{Metadata: utils.EcosystemMetadata{DID: "did:example:123"}}
	body, _ := json.Marshal(eco)
	req, _ := http.NewRequest("POST", "/ecosystems", bytes.NewBuffer(body))
	rec := httptest.NewRecorder()

	handler.CreateEcosystem(rec, req)

	assert.Equal(t, http.StatusCreated, rec.Code)
	assert.JSONEq(t, `{"message": "ecosystem created"}`, rec.Body.String())
}

func TestHandlerCreateEcosystem_BadRequest(t *testing.T) {
	svc := newTestTrustRegistryService()
	handler := NewTrustRegistryHandlers(svc)

	req, _ := http.NewRequest("POST", "/ecosystems", bytes.NewBuffer([]byte("{invalid json}")))
	rec := httptest.NewRecorder()

	handler.CreateEcosystem(rec, req)

	assert.Equal(t, http.StatusBadRequest, rec.Code)
}

func TestHandlerGetEcosystem_Success(t *testing.T) {
	svc := newTestTrustRegistryService()
	handler := NewTrustRegistryHandlers(svc)

	// First, create an ecosystem.
	eco := utils.Ecosystem{Metadata: utils.EcosystemMetadata{DID: "did:example:123"}}
	err := svc.CreateEcosystem(eco)
	assert.NoError(t, err)

	req, _ := http.NewRequest("GET", "/ecosystems/did:example:123", nil)
	rec := httptest.NewRecorder()

	handler.GetEcosystem(rec, req, "did:example:123")

	assert.Equal(t, http.StatusOK, rec.Code)

	var returned utils.Ecosystem
	err = json.Unmarshal(rec.Body.Bytes(), &returned)
	assert.NoError(t, err)
	assert.Equal(t, "did:example:123", returned.Metadata.DID)
}

func TestHandlerGetEcosystem_NotFound(t *testing.T) {
	svc := newTestTrustRegistryService()
	handler := NewTrustRegistryHandlers(svc)

	req, _ := http.NewRequest("GET", "/ecosystems/did:example:999", nil)
	rec := httptest.NewRecorder()

	handler.GetEcosystem(rec, req, "did:example:999")

	assert.Equal(t, http.StatusNotFound, rec.Code)
}

func TestHandlerUpdateEcosystem_Success(t *testing.T) {
	svc := newTestTrustRegistryService()
	handler := NewTrustRegistryHandlers(svc)

	// Create an ecosystem.
	eco := utils.Ecosystem{
		Metadata: utils.EcosystemMetadata{
			DID:  "did:example:123",
			Name: "Old Name",
		},
	}
	err := svc.CreateEcosystem(eco)
	assert.NoError(t, err)

	// Update the ecosystem.
	updatedEco := utils.Ecosystem{
		Metadata: utils.EcosystemMetadata{
			DID:  "did:example:123",
			Name: "New Name",
		},
	}
	body, _ := json.Marshal(updatedEco)
	req, _ := http.NewRequest("PUT", "/ecosystems/did:example:123", bytes.NewBuffer(body))
	rec := httptest.NewRecorder()

	handler.UpdateEcosystem(rec, req, "did:example:123")

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.JSONEq(t, `{"message": "ecosystem updated"}`, rec.Body.String())

	// Verify update.
	result, err := svc.GetEcosystem("did:example:123")
	assert.NoError(t, err)
	assert.Equal(t, "New Name", result.Metadata.Name)
}

func TestHandlerUpdateEcosystem_MismatchedDID(t *testing.T) {
	svc := newTestTrustRegistryService()
	handler := NewTrustRegistryHandlers(svc)

	eco := utils.Ecosystem{Metadata: utils.EcosystemMetadata{DID: "did:example:123"}}
	body, _ := json.Marshal(eco)
	req, _ := http.NewRequest("PUT", "/ecosystems/did:example:999", bytes.NewBuffer(body))
	rec := httptest.NewRecorder()

	handler.UpdateEcosystem(rec, req, "did:example:999")

	assert.Equal(t, http.StatusBadRequest, rec.Code)
}

func TestHandlerRemoveEcosystem_Success(t *testing.T) {
	svc := newTestTrustRegistryService()
	handler := NewTrustRegistryHandlers(svc)

	// Create an ecosystem.
	eco := utils.Ecosystem{Metadata: utils.EcosystemMetadata{DID: "did:example:123"}}
	err := svc.CreateEcosystem(eco)
	assert.NoError(t, err)

	req, _ := http.NewRequest("DELETE", "/ecosystems/did:example:123", nil)
	rec := httptest.NewRecorder()

	handler.RemoveEcosystem(rec, req, "did:example:123")

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.JSONEq(t, `{"message": "ecosystem removed"}`, rec.Body.String())

	// Verify removal.
	_, err = svc.GetEcosystem("did:example:123")
	assert.Error(t, err)
}

func TestHandlerRemoveEcosystem_NotFound(t *testing.T) {
	svc := newTestTrustRegistryService()
	handler := NewTrustRegistryHandlers(svc)

	req, _ := http.NewRequest("DELETE", "/ecosystems/did:example:999", nil)
	rec := httptest.NewRecorder()

	handler.RemoveEcosystem(rec, req, "did:example:999")

	assert.Equal(t, http.StatusNotFound, rec.Code)
}

func TestHandlerRecognizeEcosystem_Success(t *testing.T) {
	svc := newTestTrustRegistryService()
	handler := NewTrustRegistryHandlers(svc)

	// Create an ecosystem.
	eco := utils.Ecosystem{
		Metadata: utils.EcosystemMetadata{
			DID: "did:example:123",
			// Optionally set a default status.
			Status: utils.Status{Active: true},
		},
	}
	err := svc.CreateEcosystem(eco)
	assert.NoError(t, err)

	reqParams := admin.RecognizeEcosystemParams{
		Did:    "did:example:123",
		Egf:    "EGF123",
		Scope:  StringPointer("global"),
		Active: BoolPointer(true),
	}
	req, _ := http.NewRequest("POST", "/ecosystems/recognitions", nil)
	rec := httptest.NewRecorder()

	handler.RecognizeEcosystem(rec, req, reqParams)

	assert.Equal(t, http.StatusOK, rec.Code)
	var resp admin.RecognizeEcosystemResponse
	err = json.Unmarshal(rec.Body.Bytes(), &resp)
	assert.NoError(t, err)
	assert.NotNil(t, resp.Message)

	// Verify that the ecosystem now has one recognition entry.
	updatedEco, err := svc.GetEcosystem("did:example:123")
	assert.NoError(t, err)
	assert.Len(t, updatedEco.RecognitionEntries, 1)

	t.Run("recognize ecosystem no scope", func(t *testing.T) {

		reqParams := admin.RecognizeEcosystemParams{}
		req, _ := http.NewRequest("POST", "/ecosystems/recognitions", nil)
		rec := httptest.NewRecorder()

		handler.RecognizeEcosystem(rec, req, reqParams)
		assert.Equal(t, http.StatusBadRequest, rec.Code)
	})

}

func TestHandlerRecognizeEcosystem_MissingDID(t *testing.T) {
	svc := newTestTrustRegistryService()
	handler := NewTrustRegistryHandlers(svc)

	reqParams := admin.RecognizeEcosystemParams{
		// Missing Did field.
		Egf: "EGF123",
	}
	req, _ := http.NewRequest("POST", "/ecosystems/recognitions", nil)
	rec := httptest.NewRecorder()

	handler.RecognizeEcosystem(rec, req, reqParams)

	assert.Equal(t, http.StatusBadRequest, rec.Code)
}

func TestHandlerAuthorizeEntry_Success(t *testing.T) {
	svc := newTestTrustRegistryService()
	handler := NewTrustRegistryHandlers(svc)

	// Create an ecosystem.
	eco := utils.Ecosystem{Metadata: utils.EcosystemMetadata{DID: "did:example:123"}}
	err := svc.CreateEcosystem(eco)
	assert.NoError(t, err)

	// Create a request URL with query parameters.
	req, _ := http.NewRequest("POST", "/ecosystems/authorizations", nil)
	q := req.URL.Query()
	q.Set("did", "did:example:123")
	q.Set("egf", "EGF123")
	q.Set("authorization_id", "auth-456")
	q.Set("active", "true")
	req.URL.RawQuery = q.Encode()
	rec := httptest.NewRecorder()

	// For AuthorizeEntry, the handler extracts parameters from the query.
	handler.AuthorizeEntry(rec, req, admin.AuthorizeEntryParams{})

	assert.Equal(t, http.StatusOK, rec.Code)
	var authResp admin.AuthorizeEntryResponse
	err = json.Unmarshal(rec.Body.Bytes(), &authResp)
	assert.NoError(t, err)
	assert.NotNil(t, authResp.Message)

	// Verify that the ecosystem now has one authorization entry.
	updatedEco, err := svc.GetEcosystem("did:example:123")
	assert.NoError(t, err)
	assert.Len(t, updatedEco.AuthorizationEntries, 1)
}

func TestHandlerAuthorizeEntry_MissingDID(t *testing.T) {
	svc := newTestTrustRegistryService()
	handler := NewTrustRegistryHandlers(svc)

	req, _ := http.NewRequest("POST", "/ecosystems/authorizations", nil)
	// Missing required query parameter "did".
	req.URL.RawQuery = "egf=EGF123&authorization_id=auth-456"
	rec := httptest.NewRecorder()

	handler.AuthorizeEntry(rec, req, admin.AuthorizeEntryParams{})

	assert.Equal(t, http.StatusBadRequest, rec.Code)
}

func TestHandlerListEcosystems_Success(t *testing.T) {
	svc := newTestTrustRegistryService()
	// Pre-populate the registry with some ecosystems.
	ecos := []utils.Ecosystem{
		{Metadata: utils.EcosystemMetadata{DID: "did:example:123"}},
		{Metadata: utils.EcosystemMetadata{DID: "did:example:456"}},
	}
	for _, eco := range ecos {
		err := svc.CreateEcosystem(eco)
		assert.NoError(t, err)
	}

	handler := NewTrustRegistryHandlers(svc)
	req, _ := http.NewRequest("GET", "/ecosystems", nil)
	rec := httptest.NewRecorder()

	handler.ListEcosystems(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.JSONEq(t, `["did:example:123", "did:example:456"]`, rec.Body.String())
}

package registry

import (
	"testing"

	"github.com/andorsk/tswg-trust-registry-protocol/reference_implementation/gen/admin"
	"github.com/andorsk/tswg-trust-registry-protocol/reference_implementation/pkg/utils"
	"github.com/stretchr/testify/assert"
)

// --- Helper functions for pointer values ---
func newTestRegistry() *utils.TrustRegistry {
	meta := utils.Metadata{
		Identifier:    "test-registry",
		Name:          "Test Registry",
		Version:       "1.0.0",
		PrimaryEGFURI: "",
		Description:   "A registry for testing",
		Language:      "en",
	}
	return utils.NewTrustRegistry(meta)
}

func newTestTrustRegistryService() *TrustRegistryService {
	reg := newTestRegistry()
	return &TrustRegistryService{Registry: reg}
}

func TestCreateEcosystem_Success(t *testing.T) {
	service := newTestTrustRegistryService()
	newEco := utils.Ecosystem{Metadata: utils.EcosystemMetadata{DID: "did:example:123"}}
	err := service.CreateEcosystem(newEco)
	assert.NoError(t, err)
	// Verify that the registry now has one ecosystem.
	eco, err := service.GetEcosystem("did:example:123")
	assert.NoError(t, err)
	assert.Equal(t, "did:example:123", eco.Metadata.DID)
}

func TestCreateEcosystem_AlreadyExists(t *testing.T) {
	service := newTestTrustRegistryService()

	newEco := utils.Ecosystem{Metadata: utils.EcosystemMetadata{DID: "did:example:123"}}
	err := service.CreateEcosystem(newEco)
	assert.NoError(t, err)

	// Try to add the same ecosystem again.
	err = service.CreateEcosystem(newEco)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "ecosystem already exists")
}

func TestGetEcosystem_NotFound(t *testing.T) {
	service := newTestTrustRegistryService()

	_, err := service.GetEcosystem("did:example:999")
	assert.Error(t, err)
	assert.Equal(t, "ecosystem not found", err.Error())
}

func TestUpdateEcosystem_Success(t *testing.T) {
	service := newTestTrustRegistryService()

	// Create an ecosystem.
	eco := utils.Ecosystem{
		Metadata: utils.EcosystemMetadata{
			DID:  "did:example:123",
			Name: "Old Name",
		},
	}
	err := service.CreateEcosystem(eco)
	assert.NoError(t, err)

	// Update the ecosystem.
	updatedEco := utils.Ecosystem{
		Metadata: utils.EcosystemMetadata{
			DID:  "did:example:123",
			Name: "New Name",
		},
	}
	err = service.UpdateEcosystem(updatedEco)
	assert.NoError(t, err)

	// Retrieve and verify the update.
	result, err := service.GetEcosystem("did:example:123")
	assert.NoError(t, err)
	assert.Equal(t, "New Name", result.Metadata.Name)
}

func TestUpdateEcosystem_NotFound(t *testing.T) {
	service := newTestTrustRegistryService()

	// Attempt to update an ecosystem that does not exist.
	updatedEco := utils.Ecosystem{
		Metadata: utils.EcosystemMetadata{
			DID: "did:example:999",
		},
	}
	err := service.UpdateEcosystem(updatedEco)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "ecosystem not found for update")
}

func TestRemoveEcosystem_Success(t *testing.T) {
	service := newTestTrustRegistryService()

	// Create an ecosystem.
	eco := utils.Ecosystem{Metadata: utils.EcosystemMetadata{DID: "did:example:123"}}
	err := service.CreateEcosystem(eco)
	assert.NoError(t, err)

	// Remove it.
	err = service.RemoveEcosystem("did:example:123")
	assert.NoError(t, err)

	// Try to get the removed ecosystem.
	_, err = service.GetEcosystem("did:example:123")
	assert.Error(t, err)
}

func TestRemoveEcosystem_NotFound(t *testing.T) {
	service := newTestTrustRegistryService()

	err := service.RemoveEcosystem("did:example:999")
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "ecosystem not found for removal")
}

func TestRecognizeEcosystem_Success(t *testing.T) {
	service := newTestTrustRegistryService()

	// Create an ecosystem first.
	eco := utils.Ecosystem{
		Metadata: utils.EcosystemMetadata{
			DID: "did:example:123",
			// Set a default status (active) if needed.
			Status: utils.Status{Active: true},
		},
	}
	err := service.CreateEcosystem(eco)
	assert.NoError(t, err)

	req := admin.RecognizeEcosystemParams{
		Did:    "did:example:123",
		Egf:    "EGF123",
		Scope:  StringPointer("global"),
		Active: BoolPointer(true),
	}

	resp, err := service.RecognizeEcosystem(req)
	assert.NoError(t, err)
	assert.NotNil(t, resp.Message)
	assert.Equal(t, "Ecosystem recognized successfully", *resp.Message)

	// Verify that the ecosystem now has one recognition entry.
	updatedEco, err := service.GetEcosystem("did:example:123")
	assert.NoError(t, err)
	assert.Len(t, updatedEco.RecognitionEntries, 1)
	// Optionally, check details on the new recognition entry.
}

func TestRecognizeEcosystem_MissingDID(t *testing.T) {
	service := newTestTrustRegistryService()

	req := admin.RecognizeEcosystemParams{
		// Missing Did field.
		Egf: "EGF123",
	}
	_, err := service.RecognizeEcosystem(req)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "ecosystem DID is required")
}

func TestAuthorizeEntry_Success(t *testing.T) {
	service := newTestTrustRegistryService()

	// Create an ecosystem.
	eco := utils.Ecosystem{
		Metadata: utils.EcosystemMetadata{DID: "did:example:123"},
	}
	err := service.CreateEcosystem(eco)
	assert.NoError(t, err)

	req := admin.AuthorizeEntryParams{
		Did:             "did:example:123",
		Egf:             "EGF123",
		AuthorizationId: "auth-456",
		Active:          BoolPointer(true),
	}

	resp, err := service.AuthorizeEntry(req)
	assert.NoError(t, err)
	assert.NotNil(t, resp.Message)
	assert.Equal(t, "Entry authorized successfully", *resp.Message)

	// Verify that the ecosystem now has one authorization entry.
	updatedEco, err := service.GetEcosystem("did:example:123")
	assert.NoError(t, err)
	assert.Len(t, updatedEco.AuthorizationEntries, 1)
}

func TestAuthorizeEntry_MissingDID(t *testing.T) {
	service := newTestTrustRegistryService()

	req := admin.AuthorizeEntryParams{
		// Missing Did.
		AuthorizationId: "auth-456",
	}
	_, err := service.AuthorizeEntry(req)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "ecosystem DID is required")
}

func TestAuthorizeEntry_MissingAuthID(t *testing.T) {
	service := newTestTrustRegistryService()

	req := admin.AuthorizeEntryParams{
		Did: "did:example:123",
		// Missing AuthorizationId.
	}
	_, err := service.AuthorizeEntry(req)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "authorization_id is required")
}

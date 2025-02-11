package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Test GetEcosystemByDID
func TestGetEcosystemByDID_Success(t *testing.T) {
	tr := TrustRegistry{
		Ecosystems: []Ecosystem{
			{Metadata: EcosystemMetadata{DID: "did:example:123"}},
		},
	}

	eco, err := tr.GetEcosystemByDID("did:example:123")
	assert.NoError(t, err)
	assert.Equal(t, "did:example:123", eco.Metadata.DID)
}

func TestGetEcosystemByDID_NotFound(t *testing.T) {
	tr := TrustRegistry{}

	_, err := tr.GetEcosystemByDID("did:example:999")
	assert.Error(t, err)
	assert.Equal(t, "ecosystem not found", err.Error())
}

// Test GetAuthorizationTypeByKey
func TestGetAuthorizationTypeByKey_Success(t *testing.T) {
	eco := Ecosystem{
		AuthorizationTypes: map[string]AuthorizationType{
			"auth-1": {Name: "Type 1"},
		},
	}

	auth, key, err := eco.GetAuthorizationTypeByKey("auth-1")
	assert.NoError(t, err)
	assert.Equal(t, "Type 1", auth.Name)
	assert.Equal(t, "auth-1", key)
}

func TestGetAuthorizationTypeByKey_NotFound(t *testing.T) {
	eco := Ecosystem{}

	_, _, err := eco.GetAuthorizationTypeByKey("auth-unknown")
	assert.Error(t, err)
	assert.Equal(t, "authorization type not found", err.Error())
}

// Test FindRecognitionsByScope
func TestFindRecognitionsByScope_Success(t *testing.T) {
	eco := Ecosystem{
		RecognitionEntries: []RecognitionEntry{
			{Scope: "scope-1", DID: "did:example:111"},
			{Scope: "scope-1", DID: "did:example:222"},
		},
	}

	results := eco.FindRecognitionsByScope("scope-1")
	assert.Len(t, results, 2)
}

func TestFindRecognitionsByScope_NoResults(t *testing.T) {
	eco := Ecosystem{}

	results := eco.FindRecognitionsByScope("scope-missing")
	assert.Empty(t, results)
}

// Test FindAuthorizationsByScope
func TestFindAuthorizationsByScope_Success(t *testing.T) {
	eco := Ecosystem{
		AuthorizationEntries: []AuthorizationEntry{
			{Authorization: "auth-1", DID: "did:example:111"},
			{Authorization: "auth-1", DID: "did:example:222"},
		},
	}

	results := eco.FindAuthorizationsByScope("auth-1")
	assert.Len(t, results, 2)
}

func TestFindAuthorizationsByScope_NoResults(t *testing.T) {
	eco := Ecosystem{}

	results := eco.FindAuthorizationsByScope("auth-missing")
	assert.Empty(t, results)
}

// Test AddEcosystem
func TestAddEcosystem_Success(t *testing.T) {
	tr := TrustRegistry{}

	newEco := Ecosystem{Metadata: EcosystemMetadata{DID: "did:example:123"}}
	err := tr.AddEcosystem(newEco)

	assert.NoError(t, err)
	assert.Len(t, tr.Ecosystems, 1)
}

func TestAddEcosystem_AlreadyExists(t *testing.T) {
	tr := TrustRegistry{
		Ecosystems: []Ecosystem{
			{Metadata: EcosystemMetadata{DID: "did:example:123"}},
		},
	}

	newEco := Ecosystem{Metadata: EcosystemMetadata{DID: "did:example:123"}}
	err := tr.AddEcosystem(newEco)

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "ecosystem already exists")
}

// Test UpdateEcosystem
func TestUpdateEcosystem_Success(t *testing.T) {
	tr := TrustRegistry{
		Ecosystems: []Ecosystem{
			{Metadata: EcosystemMetadata{DID: "did:example:123", Name: "Old Name"}},
		},
	}

	updatedEco := Ecosystem{Metadata: EcosystemMetadata{DID: "did:example:123", Name: "New Name"}}
	err := tr.UpdateEcosystem(updatedEco)

	assert.NoError(t, err)
	assert.Equal(t, "New Name", tr.Ecosystems[0].Metadata.Name)
}

func TestUpdateEcosystem_NotFound(t *testing.T) {
	tr := TrustRegistry{}

	updatedEco := Ecosystem{Metadata: EcosystemMetadata{DID: "did:example:999"}}
	err := tr.UpdateEcosystem(updatedEco)

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "ecosystem not found for update")
}

// Test RemoveEcosystem
func TestRemoveEcosystem_Success(t *testing.T) {
	tr := TrustRegistry{
		Ecosystems: []Ecosystem{
			{Metadata: EcosystemMetadata{DID: "did:example:123"}},
		},
	}

	err := tr.RemoveEcosystem("did:example:123")

	assert.NoError(t, err)
	assert.Empty(t, tr.Ecosystems)
}

func TestRemoveEcosystem_NotFound(t *testing.T) {
	tr := TrustRegistry{}

	err := tr.RemoveEcosystem("did:example:999")

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "ecosystem not found for removal")
}

// Test Empty Cases for TrustRegistry
func TestTrustRegistry_Empty(t *testing.T) {
	tr := TrustRegistry{}

	_, err := tr.GetEcosystemByDID("did:example:unknown")
	assert.Error(t, err)

	err = tr.RemoveEcosystem("did:example:unknown")
	assert.Error(t, err)
}

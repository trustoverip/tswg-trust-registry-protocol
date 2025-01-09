package pkg

import "testing"

func TestLoadTrustRegistryFromFile(t *testing.T) {
	registry, err := LoadTrustRegistryFromFile("../data/registry.json")
	if err != nil {
		t.Fatalf("LoadTrustRegistryFromFile returned an error: %v", err)
	}
	if registry == nil {
		t.Fatal("Expected registry to be non-nil")
	}
	if registry.Metadata.Name != "GAN Trust Registry" {
		t.Errorf("Expected registry.Name to be 'TestRegistry', got '%s'", registry.Metadata.Name)
	}
}

func TestGetEcosystemByEGFURI(t *testing.T) {
	registry := TrustRegistry{
		Ecosystems: []Ecosystem{
			{
				Metadata: EcosystemMetadata{
					EGFURI: "test-egf-uri",
				},
			},
		},
	}
	ecosystem, err := registry.GetEcosystemByEGFURI("test-egf-uri")
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if ecosystem.Metadata.EGFURI != "test-egf-uri" {
		t.Errorf("Expected ecosystem with EGFURI 'test-egf-uri', got '%s'", ecosystem.Metadata.EGFURI)
	}

	_, err = registry.GetEcosystemByEGFURI("non-existent-uri")
	if err == nil {
		t.Errorf("Expected error when looking for a non-existent ecosystem")
	}
}

func TestGetEntryByVID(t *testing.T) {
	registry := TrustRegistry{
		Ecosystems: []Ecosystem{
			{
				Entries: []Entry{
					{
						DID: "test-did-1",
					},
					{
						DID: "test-did-2",
					},
				},
			},
		},
	}

	entry, err := registry.GetEntryByVID("test-did-1")
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if entry.DID != "test-did-1" {
		t.Errorf("Expected entry with DID 'test-did-1', got '%s'", entry.DID)
	}

	_, err = registry.GetEntryByVID("non-existent-did")
	if err == nil {
		t.Errorf("Expected error when looking for a non-existent entry")
	}
}

func TestGetAuthorizationTypeByKey(t *testing.T) {
	ecosystem := Ecosystem{
		AuthorizationTypes: map[string]AuthorizationType{
			"auth-key-1": {
				Name:        "Authorization 1",
				Description: "Test authorization 1",
			},
		},
	}

	authType, key, err := ecosystem.GetAuthorizationTypeByKey("auth-key-1")
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if key != "auth-key-1" {
		t.Errorf("Expected key 'auth-key-1', got '%s'", key)
	}
	if authType.Name != "Authorization 1" {
		t.Errorf("Expected authorization name 'Authorization 1', got '%s'", authType.Name)
	}

	_, _, err = ecosystem.GetAuthorizationTypeByKey("non-existent-key")
	if err == nil {
		t.Errorf("Expected error when looking for a non-existent authorization type")
	}
}

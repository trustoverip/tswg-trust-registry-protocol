package utils

import (
	"errors"
	"time"
)

func NewTrustRegistry(meta Metadata) *TrustRegistry {
	return &TrustRegistry{
		Metadata:   meta,
		Ecosystems: []Ecosystem{},
	}
}

type Ecosystem struct {
	Metadata             EcosystemMetadata            `json:"metadata"`
	Entries              []Entry                      `json:"entries"`
	AssuranceLevels      map[string]AssuranceLevel    `json:"assurance_levels"`
	AuthorizationTypes   map[string]AuthorizationType `json:"authorization_types"`
	VIDMethods           []VIDMethod                  `json:"vid_methods"`
	RecognitionEntries   []RecognitionEntry           `json:"recognition_entries,omitempty"`
	AuthorizationEntries []AuthorizationEntry         `json:"authorization_entries,omitempty"`
}

func NewEcosystem(meta EcosystemMetadata) *Ecosystem {
	return &Ecosystem{
		Metadata:             meta,
		Entries:              make([]Entry, 0),
		AssuranceLevels:      make(map[string]AssuranceLevel),
		AuthorizationTypes:   make(map[string]AuthorizationType),
		VIDMethods:           make([]VIDMethod, 0),
		RecognitionEntries:   make([]RecognitionEntry, 0),
		AuthorizationEntries: make([]AuthorizationEntry, 0),
	}
}

// EntryType defines the type of entry.
type EntryType string

const (
	RecognitionEntryType   EntryType = "recognition"
	AuthorizationEntryType EntryType = "authorization"
)

// TrustRegistry represents the overall structure of the registry.
type TrustRegistry struct {
	Metadata   Metadata    `json:"metadata"`
	Ecosystems []Ecosystem `json:"ecosystems"`
}

// GetEcosystemByDID finds an ecosystem by its DID.
func (t *TrustRegistry) GetEcosystemByDID(did string) (Ecosystem, error) {
	for _, e := range t.Ecosystems {
		if e.Metadata.DID == did {
			return e, nil
		}
	}
	return Ecosystem{}, errors.New("ecosystem not found")
}

// GetAuthorizationTypeByKey fetches an AuthorizationType by key from this ecosystem.
func (e *Ecosystem) GetAuthorizationTypeByKey(key string) (AuthorizationType, string, error) {
	for k, auth := range e.AuthorizationTypes {
		if k == key {
			return auth, k, nil
		}
	}
	return AuthorizationType{}, "", errors.New("authorization type not found")
}

// FindRecognitionsByScope returns all RecognitionEntries in this ecosystem for a given scope.
func (e *Ecosystem) FindRecognitionsByScope(scope string) []RecognitionEntry {
	var results []RecognitionEntry
	for _, rec := range e.RecognitionEntries {
		if rec.Scope == scope {
			results = append(results, rec)
		}
	}
	return results
}

// FindAuthorizationsByScope returns all AuthorizationEntries in this ecosystem for a given scope.
func (e *Ecosystem) FindAuthorizationsByScope(scope string) []AuthorizationEntry {
	var results []AuthorizationEntry
	for _, auth := range e.AuthorizationEntries {
		if auth.Authorization == scope {
			results = append(results, auth)
		}
	}
	return results
}

// Metadata defines the metadata information at the registry (top) level.
type Metadata struct {
	Identifier    string `json:"identifier"`
	Name          string `json:"name"`
	Version       string `json:"version,omitempty"`
	PrimaryEGFURI string `json:"primaryEGFURI,omitempty"`
	Description   string `json:"description,omitempty"`
	Language      string `json:"language,omitempty"`
}

// EcosystemMetadata defines the metadata of an ecosystem.
type EcosystemMetadata struct {
	Name                 string     `json:"name"`
	Type                 string     `json:"type"`
	Namespaces           []string   `json:"namespaces,omitempty"`
	Description          string     `json:"description"`
	DID                  string     `json:"did"`
	EGFURI               string     `json:"egfURI,omitempty"`
	PrimaryTrustRegistry string     `json:"primary_trust_registry,omitempty"`
	Registries           []Registry `json:"registries,omitempty"`
	Status               Status     `json:"status"`
}

// Status defines the status of an entity.
type Status struct {
	Active bool   `json:"active"`
	Detail string `json:"detail,omitempty"`
}

// Registry defines a trust registry within an ecosystem's metadata.
type Registry struct {
	Name          string `json:"name"`
	Identifier    string `json:"identifier"`
	Description   string `json:"description"`
	PrimaryEGFURI string `json:"primaryEGFURI"`
	PeerType      string `json:"peerType"`
}

// Entry defines an entry in the ecosystem.
type Entry struct {
	DID            string    `json:"did"`
	ValidFrom      time.Time `json:"validFrom,omitempty"`
	ValidTo        time.Time `json:"validTo,omitempty"`
	Status         Status    `json:"status"`
	Authorizations []string  `json:"authorizations,omitempty"`
}

// AssuranceLevel defines the assurance level of an ecosystem.
type AssuranceLevel struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

// AuthorizationType defines the authorization type of an ecosystem.
type AuthorizationType struct {
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
}

// VIDMethod defines the DID methods supported in the ecosystem.
type VIDMethod struct {
	Identifier              string `json:"identifier"`
	MaximumLevelOfAssurance string `json:"maxiumum_level_of_assurance"`
}

// RecognitionEntry captures a "Recognized" relationship in an ecosystem.
type RecognitionEntry struct {
	Scope  string `json:"scope"`
	DID    string `json:"did"`
	Status Status `json:"status"`
}

// AuthorizationEntry captures an "Authorized" relationship in an ecosystem.
// Note that the JSON now uses the key "authorization" to indicate the type.
type AuthorizationEntry struct {
	Authorization string `json:"authorization"`
	ID            string `json:"id"`
	DID           string `json:"did"`
	// The status field is now optional. In its absence, we assume "active."
	Status Status `json:"status,omitempty"`
}

func (t *TrustRegistry) AddEcosystem(newEco Ecosystem) error {
	for _, e := range t.Ecosystems {
		if e.Metadata.DID == newEco.Metadata.DID {
			return errors.New("ecosystem already exists with DID: " + newEco.Metadata.DID)
		}
	}
	t.Ecosystems = append(t.Ecosystems, newEco)
	return nil
}

// UpdateEcosystem replaces an ecosystem with a matching DID.
func (t *TrustRegistry) UpdateEcosystem(updatedEco Ecosystem) error {
	for i, e := range t.Ecosystems {
		if e.Metadata.DID == updatedEco.Metadata.DID {
			t.Ecosystems[i] = updatedEco
			return nil
		}
	}
	return errors.New("ecosystem not found for update: " + updatedEco.Metadata.DID)
}

// RemoveEcosystem removes an ecosystem by DID.
func (t *TrustRegistry) RemoveEcosystem(did string) error {
	for i, e := range t.Ecosystems {
		if e.Metadata.DID == did {
			t.Ecosystems = append(t.Ecosystems[:i], t.Ecosystems[i+1:]...)
			return nil
		}
	}
	return errors.New("ecosystem not found for removal: " + did)
}

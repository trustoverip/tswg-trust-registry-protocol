package pkg

import (
	"errors"
	"time"
)

// TrustRegistry represents the overall structure of the registry
type TrustRegistry struct {
	Metadata      Metadata       `json:"metadata"`
	Organizations []Organization `json:"organizations"`
	Ecosystems    []Ecosystem    `json:"ecosystems"`
}

func (t *TrustRegistry) GetEcosystemByEGFURI(EGFURI string) (Ecosystem, error) {
	for _, e := range t.Ecosystems {
		if e.Metadata.EGFURI == EGFURI {
			return e, nil
		}
	}
	return Ecosystem{}, errors.New("ecosystem not found")
}

func (r *TrustRegistry) GetEntryByVID(VID string) (Entry, error) {
	for _, e := range r.Ecosystems {
		for _, entry := range e.Entries {
			if entry.DID == VID {
				return entry, nil
			}
		}
	}
	return Entry{}, errors.New("entry not found")
}

// Metadata defines the metadata information
type Metadata struct {
	Identifier    string `json:"identifier"`
	Name          string `json:"name"`
	Version       string `json:"version"`
	PrimaryEGFURI string `json:"primaryEGFURI"`
	Description   string `json:"description"`
	Language      string `json:"language"`
}

// Organization defines the structure of an organization
type Organization struct {
	Name                    string   `json:"name"`
	Type                    string   `json:"type"`
	ParticipatingNamespaces []string `json:"participatingNamespaces,omitempty"`
	DID                     string   `json:"did"`
	Status                  Status   `json:"status"`
}

// Ecosystem defines the structure of an ecosystem
type Ecosystem struct {
	Metadata           EcosystemMetadata            `json:"metadata"`
	Entries            []Entry                      `json:"entries"`
	AssuranceLevels    map[string]AssuranceLevel    `json:"assurance_levels"`
	AuthorizationTypes map[string]AuthorizationType `json:"authorization_types"`
	VIDMethods         []VIDMethod                  `json:"vid_methods"`
}

func (e *Ecosystem) GetAuthorizationTypeByKey(key string) (auth AuthorizationType, id string, err error) {
	for k, auth := range e.AuthorizationTypes {
		if k == key {
			return auth, key, nil
		}
	}
	return AuthorizationType{}, "", errors.New("authorization not found")
}

// EcosystemMetadata defines the metadata of an ecosystem
type EcosystemMetadata struct {
	Name                 string     `json:"name"`
	Type                 string     `json:"type"`
	Namespaces           []string   `json:"namespaces"`
	Description          string     `json:"description"`
	DID                  string     `json:"did"`
	EGFURI               string     `json:"egfURI"`
	PrimaryTrustRegistry string     `json:"primary_trust_registry"`
	Registries           []Registry `json:"registries"`
	Status               Status     `json:"status"`
}

// Status defines the status of an entity
type Status struct {
	Status string `json:"status"`
	Detail string `json:"detail"`
}

// Registry defines a trust registry
type Registry struct {
	Name          string `json:"name"`
	Identifier    string `json:"identifier"`
	Description   string `json:"description"`
	PrimaryEGFURI string `json:"primaryEGFURI"`
	PeerType      string `json:"peerType"`
}

// Entry defines an entry in the ecosystem
type Entry struct {
	DID            string    `json:"did"`
	ValidFrom      time.Time `json:"validFrom"`
	ValidTo        time.Time `json:"validTo"`
	Status         Status    `json:"status"`
	Authorizations []string  `json:"authorizations,omitempty"`
}

// AssuranceLevel defines the assurance level of an ecosystem
type AssuranceLevel struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

// AuthorizationType defines the authorization type of an ecosystem
type AuthorizationType struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

// VIDMethod defines the DID methods supported in the ecosystem
type VIDMethod struct {
	Identifier              string `json:"identifier"`
	MaximumLevelOfAssurance string `json:"maxiumum_level_of_assurance"`
}

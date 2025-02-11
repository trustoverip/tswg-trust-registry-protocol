package registry

import (
	"errors"
	"fmt"

	"github.com/andorsk/tswg-trust-registry-protocol/reference_implementation/gen/admin"
	"github.com/andorsk/tswg-trust-registry-protocol/reference_implementation/pkg/utils"
)

type TrustRegistryServiceInterface interface {
	CreateEcosystem(eco utils.Ecosystem) error
	GetEcosystem(did string) (utils.Ecosystem, error)
	UpdateEcosystem(eco utils.Ecosystem) error
	RemoveEcosystem(did string) error
	RecognizeEcosystem(req admin.RecognizeEcosystemParams) (admin.RecognizeEcosystemResponse, error)
	AuthorizeEntry(req admin.AuthorizeEntryParams) (admin.AuthorizeEntryResponse, error)
}

// TrustRegistryService is a struct that wraps a TrustRegistry instance.
type TrustRegistryService struct {
	Registry *utils.TrustRegistry
}

// CreateEcosystem adds a new ecosystem to the registry.
func (svc *TrustRegistryService) CreateEcosystem(eco utils.Ecosystem) error {
	return svc.Registry.AddEcosystem(eco)
}

// GetEcosystem retrieves an ecosystem by DID.
func (svc *TrustRegistryService) GetEcosystem(did string) (utils.Ecosystem, error) {
	return svc.Registry.GetEcosystemByDID(did)
}

// UpdateEcosystem updates an existing ecosystem in the registry.
func (svc *TrustRegistryService) UpdateEcosystem(eco utils.Ecosystem) error {
	return svc.Registry.UpdateEcosystem(eco)
}

// RemoveEcosystem removes an existing ecosystem by DID.
func (svc *TrustRegistryService) RemoveEcosystem(did string) error {
	return svc.Registry.RemoveEcosystem(did)
}

// RecognizeEcosystem applies a recognition to an ecosystem using admin.RecognizeEcosystemParams,
// returning an admin.RecognizeEcosystemResponse if successful.
func (svc *TrustRegistryService) RecognizeEcosystem(
	req admin.RecognizeEcosystemParams,
) (admin.RecognizeEcosystemResponse, error) {

	// Make sure we have a DID to look up.
	if req.Did == "" {
		return admin.RecognizeEcosystemResponse{},
			errors.New("ecosystem DID is required in query param 'did'")
	}

	// Look up the ecosystem by DID
	eco, err := svc.Registry.GetEcosystemByDID(req.Egf)
	if err != nil {
		return admin.RecognizeEcosystemResponse{},
			fmt.Errorf("ecosystem not found: %w", err)
	}

	// Build the new recognition entry
	newEntry := utils.RecognitionEntry{
		DID: req.Did, // Ecosystem DID or recognized DID
		Status: utils.Status{
			Active: true, // Default
		},
	}

	egfVal := req.Egf
	if egfVal != "" {
		newEntry.Status.Detail = "Recognized by EGF: " + egfVal
	}

	// Append to the ecosystem's recognition entries
	eco.RecognitionEntries = append(eco.RecognitionEntries, newEntry)

	// Save updated ecosystem
	if err := svc.Registry.UpdateEcosystem(eco); err != nil {
		return admin.RecognizeEcosystemResponse{}, err
	}

	var msg = "Ecosystem recognized successfully"
	return admin.RecognizeEcosystemResponse{
		Message: &msg,
	}, nil
}

// AuthorizeEntry authorizes an ecosystem entry using admin.AuthorizeEntryParams,
// returning an admin.AuthorizeEntryResponse if successful.
func (svc *TrustRegistryService) AuthorizeEntry(
	req admin.AuthorizeEntryParams,
) (admin.AuthorizeEntryResponse, error) {

	// Make sure we have a DID to look up.
	if req.Did == "" {
		return admin.AuthorizeEntryResponse{},
			fmt.Errorf("ecosystem DID is required in query param 'did'")
	}

	// Locate the ecosystem by DID
	eco, err := svc.Registry.GetEcosystemByDID(req.Egf)
	if err != nil {
		return admin.AuthorizeEntryResponse{},
			fmt.Errorf("ecosystem not found: %w", err)
	}

	// Make sure we have an AuthorizationId
	if req.AuthorizationId == "" {
		return admin.AuthorizeEntryResponse{},
			fmt.Errorf("authorization_id is required")
	}

	// Prepare a new authorization entry
	newAuth := utils.AuthorizationEntry{
		ID:            req.AuthorizationId,
		DID:           req.Did, // or the "authorized subject" DID
		Authorization: req.Egf, // e.g., your 'egf' is controlling or naming the authorization
		Status: utils.Status{
			Active: true,
		},
	}

	// If Active is provided, set it
	if req.Active != nil {
		newAuth.Status.Active = *req.Active
	}

	// Append to the ecosystem's authorization entries
	eco.AuthorizationEntries = append(eco.AuthorizationEntries, newAuth)

	// Save updated ecosystem
	if err := svc.Registry.UpdateEcosystem(eco); err != nil {
		return admin.AuthorizeEntryResponse{}, err
	}

	msg := "Entry authorized successfully"
	return admin.AuthorizeEntryResponse{
		Message: &msg,
	}, nil
}

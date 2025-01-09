package pkg

import (
	"encoding/json"
	"net/http"

	"github.com/GANfoundation/m/v2/demo/tr-demo/gen/trqp"
)

// TRQPHandler implements the ServerInterface
type TRQPHandler struct {
	registry *TrustRegistry
}

func NewTRQPHandler(registry *TrustRegistry) *TRQPHandler {
	return &TRQPHandler{
		registry: registry,
	}
}

// GetEntitiesEntityVIDAuthorization returns authorization for an entity
func (h TRQPHandler) GetEntitiesEntityVIDAuthorization(w http.ResponseWriter, r *http.Request, entityVID trqp.VID, params trqp.GetEntitiesEntityVIDAuthorizationParams) {
	entity, err := h.registry.GetEntryByVID(entityVID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	var ret trqp.AuthorizationListType
	for _, k := range entity.Authorizations {
		auth, key, err := h.registry.Ecosystems[0].GetAuthorizationTypeByKey(k) // FIX: API needs to specify governance EGF
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		if params.AuthorizationVID != nil {
			if *params.AuthorizationVID == key {
				ret = append(ret, trqp.AuthorizationType{
					Description: auth.Description,
					Identifier:  key,
					Simplename:  auth.Name,
				})
			}
		} else {
			ret = append(ret, trqp.AuthorizationType{
				Description: auth.Description,
				Identifier:  key,
				Simplename:  auth.Name,
			})

		}

	}
	json.NewEncoder(w).Encode(ret)
}

// GetEntitiesEntityVIDAuthorizations returns authorizations for an entity
func (h TRQPHandler) GetEntitiesEntityVIDAuthorizations(w http.ResponseWriter, r *http.Request, entityVID trqp.VID) {
	entity, err := h.registry.GetEntryByVID(entityVID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	var ret trqp.AuthorizationListType
	for _, k := range entity.Authorizations {
		auth, key, err := h.registry.Ecosystems[0].GetAuthorizationTypeByKey(k) // FIX: API needs to specify governance EGF
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		ret = append(ret, trqp.AuthorizationType{
			Description: auth.Description,
			Identifier:  key,
			Simplename:  auth.Name,
		})
	}
	json.NewEncoder(w).Encode(ret)
}

// GetEntititiesEntityid returns information about a particular entity
func (h TRQPHandler) GetEntititiesEntityid(w http.ResponseWriter, r *http.Request, entityid trqp.Uri, params trqp.GetEntititiesEntityidParams) {
	entity, err := h.registry.GetEntryByVID(entityid)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	validFrom := entity.ValidFrom
	validTo := entity.ValidTo
	var statusEnum trqp.StatusTypeStatus
	switch entity.Status.Status {
	case "current":
		statusEnum = trqp.Current
	case "expired":
		statusEnum = trqp.Expired
	case "revoked":
		statusEnum = trqp.Revoked
	case "terminated":
		statusEnum = trqp.Terminated
	default:
		statusEnum = trqp.Current // Default to current if the status is unrecognized
	}
	ret := &trqp.EntityType{
		EntityVID: &entity.DID,
		EntityDataValidity: &trqp.ValidityDatesType{
			ValidFromDT:  &validFrom,
			ValidUntilDT: &validTo,
		},
		PrimaryTrustRegistryVID: &entity.DID,
		RegistrationStatus: &trqp.StatusType{
			Status: statusEnum,
			Detail: &entity.Status.Detail,
		},
		SecondaryTrustRegistries: &[]string{}, // Assuming it's an empty array for now
	}
	json.NewEncoder(w).Encode(ret)
}

// GetLookupAssurancelevels returns assurance levels
func (h TRQPHandler) GetLookupAssurancelevels(w http.ResponseWriter, r *http.Request, params trqp.GetLookupAssurancelevelsParams) {
	ecosystem, err := h.registry.GetEcosystemByEGFURI(params.EgfURI)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(ecosystem.AssuranceLevels)
}

// GetLookupAuthorizations returns authorizations
func (h TRQPHandler) GetLookupAuthorizations(w http.ResponseWriter, r *http.Request, params trqp.GetLookupAuthorizationsParams) {
	ecosystem, err := h.registry.GetEcosystemByEGFURI(params.EgfURI)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(ecosystem.AuthorizationTypes)
}

// GetLookupNamespaces returns supported namespaces
func (h TRQPHandler) GetLookupNamespaces(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(h.registry.Ecosystems[0].Metadata.Namespaces)
}

// GetLookupVidmethods returns DID methods supported by the trust registry
func (h TRQPHandler) GetLookupVidmethods(w http.ResponseWriter, r *http.Request, params trqp.GetLookupVidmethodsParams) {
	ecosystem, err := h.registry.GetEcosystemByEGFURI(params.EgfURI)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(ecosystem.VIDMethods)
}

// GetMetadata returns metadata for the trust registry
func (h TRQPHandler) GetMetadata(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(h.registry.Metadata)
}

// GetOfflineFile serves the full registry data
func (h TRQPHandler) GetOfflineFile(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(h.registry)
}

// GetTED serves a Trust Establishment Document
func (h TRQPHandler) GetTED(w http.ResponseWriter, r *http.Request) {
	// Implement your logic for TED
	http.Error(w, "Not Implemented", http.StatusNotImplemented)
}

func registryToRegistryType(r Registry) trqp.RegistryType {
	var peerType *trqp.RegistryTypePeerType
	switch r.PeerType {
	case "peer":
		peer := trqp.Peer
		peerType = &peer
	case "superior":
		superior := trqp.Superior
		peerType = &superior
	case "subordinate":
		subordinate := trqp.Subordinate
		peerType = &subordinate
	case "metaregistry":
		metaregistry := trqp.Metaregistry
		peerType = &metaregistry
	}
	return trqp.RegistryType{
		Name:                   r.Name,
		Identifier:             r.Identifier,
		Description:            &r.Description,
		PrimaryEGFURI:          &r.PrimaryEGFURI,
		PeerType:               peerType,
		AdditionalEGFURIs:      nil,
		ParticipatingNamepaces: nil,
	}
}

func (h *TRQPHandler) GetRegistriesRecognizedRegistries(w http.ResponseWriter, r *http.Request, params trqp.GetRegistriesRecognizedRegistriesParams) {
	// Initialize a slice to collect all recognized registries
	var recognizedRegistries trqp.RegistryListType

	// Filter based on NamespaceVID if provided
	if params.NamespaceVID != nil {
		for _, ecosystem := range h.registry.Ecosystems {
			for _, namespace := range ecosystem.Metadata.Namespaces {
				if namespace == *params.NamespaceVID {
					for _, registry := range ecosystem.Metadata.Registries {
						r := registryToRegistryType(registry)
						recognizedRegistries = append(recognizedRegistries, r)
					}
				}
			}
		}
	} else if params.EGFVID != nil {
		for _, ecosystem := range h.registry.Ecosystems {
			if ecosystem.Metadata.DID == *params.EGFVID {
				for _, registry := range ecosystem.Metadata.Registries {
					r := registryToRegistryType(registry)
					recognizedRegistries = append(recognizedRegistries, r)
				}
			}
		}
	} else {
		// If no NamespaceVID is provided, gather recognized registries from all ecosystems
		for _, ecosystem := range h.registry.Ecosystems {
			for _, registry := range ecosystem.Metadata.Registries {
				r := registryToRegistryType(registry)
				recognizedRegistries = append(recognizedRegistries, r)
			}

		}
	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(recognizedRegistries); err != nil {
		http.Error(w, "Failed to encode recognized registries", http.StatusInternalServerError)
		return
	}
}

// GetRegistriesRegistryVID returns data for a specific registry
func (h TRQPHandler) GetRegistriesRegistryVID(w http.ResponseWriter, r *http.Request, registryVID trqp.VID) {
	var recognizedRegistries trqp.RegistryListType
	for _, ecosystem := range h.registry.Ecosystems {
		for _, registry := range ecosystem.Metadata.Registries {
			if registry.Identifier == registryVID {
				r := registryToRegistryType(registry)
				recognizedRegistries = append(recognizedRegistries, r)
			}
		}

	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(recognizedRegistries); err != nil {
		http.Error(w, "Failed to encode recognized registries", http.StatusInternalServerError)
		return
	}

}

// GetRegistriesRegistryVIDRecognizedRegistries returns recognized registries under a specific registryVID
func (h TRQPHandler) GetRegistriesRegistryVIDRecognizedRegistries(w http.ResponseWriter, r *http.Request, registryVID trqp.VID) {
	var recognizedRegistries trqp.RegistryListType
	for _, ecosystem := range h.registry.Ecosystems {
		for _, registry := range ecosystem.Metadata.Registries {
			if registry.Identifier == registryVID {
				r := registryToRegistryType(registry)
				recognizedRegistries = append(recognizedRegistries, r)
			}
		}

	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(recognizedRegistries); err != nil {
		http.Error(w, "Failed to encode recognized registries", http.StatusInternalServerError)
		return
	}

}

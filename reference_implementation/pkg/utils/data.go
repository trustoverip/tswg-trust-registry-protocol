package utils

import (
	"encoding/json"
	"os"
)

func LoadTrustRegistryFromFile(path string) (*TrustRegistry, error) {
	var registry TrustRegistry
	dat, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(dat, &registry)
	if err != nil {
		return nil, err
	}
	return &registry, nil
}

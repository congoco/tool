package config

import (
	_ "embed"
	"encoding/json"

	"go.yaml.in/yaml/v4"
)

const customConfigFileName string = "congoco.yaml"

//go:embed default.yaml
var DefaultConfig []byte

//go:embed version.json
var VersionJson []byte

type jsonFile struct {
	Version string `json:"version"`
}

type Repository struct{}

func NewRepository() *Repository {
	r := Repository{}
	return &r
}

func (r *Repository) GetDefaults(*Parameters) (*Parameters, error) {
	params := NewParameters()

	err := yaml.Unmarshal(DefaultConfig, &params)
	if err != nil {
		return nil, err
	}

	return params, nil
}

func (r *Repository) GetVersion() (string, error) {
	jsonFile := jsonFile{}

	err := json.Unmarshal(VersionJson, &jsonFile)
	if err != nil {
		return "", err
	}
	return jsonFile.Version, nil
}

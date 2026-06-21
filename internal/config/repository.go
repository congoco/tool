package config

import (
	_ "embed"
	"encoding/json"
	"os"

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

func (r *Repository) GetVersion() (string, error) {
	jsonFile := jsonFile{}

	err := json.Unmarshal(VersionJson, &jsonFile)
	if err != nil {
		return "", err
	}
	return jsonFile.Version, nil
}

func (r *Repository) GetDefaults(params *Parameters) (*Parameters, error) {
	err := yaml.Unmarshal(DefaultConfig, params)
	if err != nil {
		return nil, err
	}

	return params, nil
}

func (r *Repository) GetCustomYaml(params *Parameters) (*Parameters, error) {
	_, err := os.Stat(customConfigFileName)
	if err != nil && os.IsNotExist(err) {
		return params, nil
	}

	if err != nil && !os.IsNotExist(err) {
		return nil, err
	}

	data, err := os.ReadFile(customConfigFileName)
	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal(data, params)
	if err != nil {
		return nil, err
	}

	return params, nil
}

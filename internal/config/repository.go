package config

import (
	_ "embed"

	"go.yaml.in/yaml/v4"
)

const customConfigFileName string = "congoco.yaml"

//go:embed default.yaml
var DefaultConfig []byte

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

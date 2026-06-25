package config

import (
	_ "embed"
	"fmt"
	"os"

	"go.yaml.in/yaml/v4"
)

//go:embed default.yaml
var DefaultConfig []byte

type Repository struct{}

func NewRepository() *Repository {
	r := Repository{}
	return &r
}

func (r *Repository) GetDefaults(cfg *Config) (*Config, error) {
	err := yaml.Unmarshal(DefaultConfig, cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}

func (r *Repository) GetCustomYaml(cfg *Config, customYamlFile bool) (*Config, error) {
	_, err := os.Stat(cfg.CustomConfigFilename)
	if err != nil && os.IsNotExist(err) {
		if customYamlFile {
			return nil, err
		}
		// fmt.Println(fmt.Errorf("No file"))
		return cfg, nil
	}

	if err != nil && !os.IsNotExist(err) {
		// fmt.Println(fmt.Errorf("Wrong file"))
		return nil, err
	}

	data, err := os.ReadFile(cfg.CustomConfigFilename)
	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal(data, cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}

func (r *Repository) SaveConfig(cfg *Config, force bool) error {
	_, err := os.Stat(cfg.CustomConfigFilename)
	if err == nil && !force {
		return fmt.Errorf("File %s already exists", cfg.CustomConfigFilename)
	}

	if err != nil && !os.IsNotExist(err) {
		return err
	}

	data, err := yaml.Marshal(cfg)
	if err != nil {
		return err
	}

	err = os.WriteFile(
		cfg.CustomConfigFilename,
		data,
		0o644,
	)
	if err != nil {
		return err
	}
	return nil
}

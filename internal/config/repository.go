package config

import (
	_ "embed"
	"os"

	"go.yaml.in/yaml/v4"
)

//go:embed default.yaml
var DefaultConfig []byte

//go:embed version.json
var VersionJson []byte

// type jsonFile struct {
// 	Version string `json:"version"`
// }

type Repository struct{}

func NewRepository() *Repository {
	r := Repository{}
	return &r
}

// func (r *Repository) GetVersion() (string, error) {
// 	jsonFile := jsonFile{}

// 	err := json.Unmarshal(VersionJson, &jsonFile)
// 	if err != nil {
// 		return "", err
// 	}
// 	return jsonFile.Version, nil
// }

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

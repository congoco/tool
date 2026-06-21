package cli

import (
	"fmt"
	"os"

	"congoco/internal/config"

	"go.yaml.in/yaml/v4"
)

type Repository struct{}

func NewRepository() (*Repository, error) {
	r := Repository{}
	return &r, nil
}

func (r *Repository) SaveConfig(params *config.Parameters, force bool) error {
	_, err := os.Stat(params.CustomConfigPath)
	if err == nil && !force {
		return fmt.Errorf("File %s already exists", params.CustomConfigPath)
	}

	if err != nil && !os.IsNotExist(err) {
		return err
	}

	data, err := yaml.Marshal(params)
	if err != nil {
		panic(err)
	}

	err = os.WriteFile(
		params.CustomConfigPath,
		data,
		0o644,
	)
	if err != nil {
		panic(err)
	}
	return nil
}

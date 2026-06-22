package congoco

import (
	_ "embed"
	"encoding/json"
)

//go:embed version.json
var VersionJson []byte

type jsonFile struct {
	Version string `json:"version"`
}

type Repository struct{}

func NewRepository() (*Repository, error) {
	r := Repository{}
	return &r, nil
}

func (r *Repository) GetVersion() (string, error) {
	jsonFile := jsonFile{}

	err := json.Unmarshal(VersionJson, &jsonFile)
	if err != nil {
		return "", err
	}
	return jsonFile.Version, nil
}

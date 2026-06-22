package congoco

import (
	_ "embed"
	"encoding/json"

	"github.com/go-git/go-git/v6"
	"github.com/go-git/go-git/v6/plumbing/object"
)

//go:embed version.json
var VersionJson []byte

type jsonFile struct {
	Version string `json:"version"`
}

type Repository struct {
	*git.Repository
}

func NewRepository() (*Repository, error) {
	repo, err := git.PlainOpenWithOptions(".", &git.PlainOpenOptions{
		DetectDotGit: true,
	})
	if err != nil {
		return nil, err
	}
	defer repo.Close()

	r := Repository{
		Repository: repo,
	}

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

func (r *Repository) GetCommits() ([]*object.Commit, error) {
	ref, err := r.Head()
	if err != nil {
		return nil, err
	}

	repoCommits, err := r.Log(&git.LogOptions{
		From: ref.Hash(),
	})
	if err != nil {
		return nil, err
	}
	commits := []*object.Commit{}
	err = repoCommits.ForEach(func(repoCommit *object.Commit) error {
		commits = append(commits, repoCommit)

		return nil
	})
	if err != nil {
		return nil, err
	}
	return commits, nil
}

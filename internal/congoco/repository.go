package congoco

import (
	_ "embed"
	"encoding/json"

	"congoco/internal/config"

	"github.com/go-git/go-git/v6"
	"github.com/go-git/go-git/v6/plumbing"
	"github.com/go-git/go-git/v6/plumbing/object"
)

//go:embed version.json
var VersionJson []byte

type jsonFile struct {
	Version string `json:"version"`
}

type Repository struct {
	cfg *config.Config
	*git.Repository
}

func NewRepository(cfg *config.Config) (*Repository, error) {
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

func (r *Repository) GetCommits(from plumbing.Hash, to plumbing.Hash) ([]*object.Commit, error) {
	repoCommits, err := r.Log(&git.LogOptions{
		From:  from,
		Order: git.LogOrderDFSPostFirstParent,
		To:    to,
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

func (r *Repository) GetTags() ([]*object.Tag, error) {
	repoTags, err := r.Tags()
	if err != nil {
		return nil, err
	}
	tags := []*object.Tag{}
	err = repoTags.ForEach(func(ref *plumbing.Reference) error {
		obj, err := r.TagObject(ref.Hash())
		switch err {
		case nil:
			// Tag object present
			tags = append(tags, obj)
		case plumbing.ErrObjectNotFound:
			// Not a tag object
		default:
			// Some other error
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return tags, nil
}

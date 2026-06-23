package congoco

import (
	"fmt"
	"slices"
	"strings"

	"congoco/internal/config"

	"github.com/go-git/go-git/v6/plumbing/object"
)

const rootPackageName string = "_root"

type CongocoRepository interface {
	GetVersion() (string, error)
	GetCommits() ([]*object.Commit, error)
	GetTags() ([]*object.Tag, error)
}

type Service struct {
	repo CongocoRepository
}

func NewService() (*Service, error) {
	repo, err := NewRepository()
	if err != nil {
		return nil, err
	}
	s := Service{
		repo: repo,
	}
	return &s, nil
}

func (s *Service) LoadVersion() (string, error) {
	version, err := s.repo.GetVersion()
	if err != nil {
		return "", err
	}
	return version, nil
}

func (s *Service) ParseMessage(message string) (*CommitMessage, error) {
	header, _, _ := strings.Cut(message, "\n")
	header = strings.TrimSpace(header)

	conv, subject, found := strings.Cut(header, ":")
	if !found {
		return nil, fmt.Errorf("No commit type")
	}

	breakingChange := false
	if strings.HasSuffix(conv, "!") {
		breakingChange = true
		conv = strings.TrimSuffix(conv, "!")
	}

	scope := ""
	typeStr, scopeStr, scopeFound := strings.Cut(conv, "(")
	if scopeFound {
		afterScope := ""
		scope, afterScope, found = strings.Cut(scopeStr, ")")
		if !found || len(afterScope) > 0 {
			return nil, fmt.Errorf("Invalid scope: %s", strings.TrimPrefix(conv, typeStr))
		}
	}

	cType, err := ParseCommitType(typeStr)
	if err != nil {
		return nil, err
	}

	cm := CommitMessage{
		BreakingChange: breakingChange,
		Scope:          scope,
		Subject:        subject,
		Type:           cType,
	}

	return &cm, nil
}

func (s *Service) ValidateBranch() ([]string, error) {
	repoCommits, err := s.repo.GetCommits()
	if err != nil {
		return nil, err
	}

	invalidCommits := make([]string, 0, len(repoCommits))
	valid := true
	for _, c := range repoCommits {
		_, err := s.ParseMessage(c.Message)
		if err != nil {
			valid = false
			invalidMessage, _, _ := strings.Cut(c.Message, "\n")
			invalidCommit := fmt.Sprintf("%s \"%s\" by %s", c.Hash.String()[:7], invalidMessage, &c.Author)
			invalidCommits = append(invalidCommits, invalidCommit)
		}
	}
	if !valid {
		return invalidCommits, fmt.Errorf("Invalid commits in branch")
	}
	return nil, nil
}

func (s *Service) GetPackageVersions(packages map[string]config.Package, cfg *config.Config) (map[string]*Version, error) {
	tags, err := s.repo.GetTags()
	if err != nil {
		return nil, err
	}

	slices.Reverse(tags)

	packages[rootPackageName] = config.Package{}
	versions := map[string]*Version{}

	for pckgName := range packages {
		for _, tag := range tags {
			pckgPrefix := pckgName
			if pckgName == rootPackageName {
				pckgPrefix = ""
			}

			version, err := s.parsePackageVersion(pckgPrefix, tag, cfg.TagPrefix)
			if err == nil {
				if pckgName == "" {
					pckgName = rootPackageName
				}
				versions[pckgName] = version
				break
			}
		}

		if versions[pckgName] == nil {
			versions[pckgName] = &Version{
				Prefix: cfg.TagPrefix,
			}
		}
	}

	return versions, nil
}

func (s *Service) parsePackageVersion(packageName string, tag *object.Tag, tagPrefix string) (*Version, error) {
	if !strings.HasPrefix(tag.Name, packageName) {
		return nil, fmt.Errorf("Not a package version tag")
	}
	str := strings.TrimPrefix(tag.Name, fmt.Sprintf("%s-", packageName))
	taggedCommit, err := tag.Commit()
	if err != nil {
		return nil, err
	}
	version, err := VersionFromString(taggedCommit.Hash.String()[:7], str, tagPrefix)
	if err != nil {
		return nil, err
	}

	return version, nil
}

func (s *Service) BuildChangelog(from, to string) (*Changelog, error) {
	cl := Changelog{}
	return &cl, nil
}

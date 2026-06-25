package congoco

import (
	"fmt"
	"slices"
	"strings"

	"congoco/internal/config"

	"github.com/go-git/go-git/v6/plumbing"
	"github.com/go-git/go-git/v6/plumbing/object"
)

type CongocoRepository interface {
	GetVersion() (string, error)
	GetCommits(from plumbing.Hash, to plumbing.Hash) ([]*object.Commit, error)
	GetTags() ([]*object.Tag, error)
	Head() (*plumbing.Reference, error)
}

type Service struct {
	cfg  *config.Config
	repo CongocoRepository
}

func NewService(cfg *config.Config) (*Service, error) {
	repo, err := NewRepository(cfg)
	if err != nil {
		return nil, err
	}
	s := Service{
		cfg:  cfg,
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
	head, err := s.repo.Head()
	if err != nil {
		return nil, err
	}

	packagesList := make([]string, 0, len(s.cfg.Packages))
	for pckgName := range s.cfg.Packages {
		packagesList = append(packagesList, pckgName)
	}

	repoCommits, err := s.repo.GetCommits(head.Hash(), plumbing.ZeroHash)
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

func (s *Service) GetPackageVersions() (map[string]*Version, error) {
	tags, err := s.repo.GetTags()
	if err != nil {
		return nil, err
	}

	slices.Reverse(tags)

	packageCurrentVersions := make(map[string]*Version, len(s.cfg.Packages))

	for pckgName := range s.cfg.Packages {
		for _, tag := range tags {
			pckgPrefix := ""
			if pckgName != s.cfg.RootPackageName {
				pckgPrefix = pckgName
			}

			version, err := s.parsePackageVersion(pckgPrefix, tag, s.cfg.TagPrefix)
			if err == nil {
				packageCurrentVersions[pckgName] = version
				break
			}
		}

		if packageCurrentVersions[pckgName] == nil {
			packageCurrentVersions[pckgName] = &Version{
				Prefix: s.cfg.TagPrefix,
			}
		}
	}

	return packageCurrentVersions, nil
}

func (s *Service) parsePackageVersion(packageName string, tag *object.Tag, tagPrefix string) (*Version, error) {
	if !strings.HasPrefix(tag.Name, packageName) {
		return nil, fmt.Errorf("Not a package version tag")
	}
	str := strings.TrimPrefix(tag.Name, fmt.Sprintf("%s-", packageName))

	version, err := ParseVersion(tag, str, tagPrefix)
	if err != nil {
		return nil, err
	}

	return version, nil
}

func (s *Service) CalculatePackagesVersions(invalidCommitsStrategy, belongCommitsStrategy string) (map[string]*Version, error) {
	packageCurrentVersions, err := s.GetPackageVersions()
	if err != nil {
		return nil, err
	}

	// packageNextVersions := make(map[string]*Version, len(s.cfg.Packages))
	// for pckgName, version := range packageCurrentVersions {
	// 	nextVersion, err := s.calculatePackageVersion(pckgName, version)
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// 	packageNextVersions[pckgName] = nextVersion
	// }
	return packageCurrentVersions, nil
}

func (s *Service) calculatePackageVersion(pckgName string, lastVersion *Version) (*Version, error) {
	head, err := s.repo.Head()
	if err != nil {
		return nil, err
	}

	lastVersionCommitHash := plumbing.ZeroHash

	if lastVersion.Tag != nil {
		taggedCommit, err := lastVersion.Tag.Commit()
		if err != nil {
			return nil, err
		}
		lastVersionCommitHash = taggedCommit.Hash
	}

	repoCommits, err := s.repo.GetCommits(head.Hash(), lastVersionCommitHash)
	if err != nil {
		return nil, err
	}

	repoCommits = repoCommits[:len(repoCommits)-1]

	/// === /// === ///

	fmt.Println(pckgName, len(repoCommits))

	commits := make([]Commit, 0, len(repoCommits))

	for _, repoCommit := range repoCommits {
		cm, err := s.ParseMessage(repoCommit.Message)
		if err != nil {
			return nil, err
		}
		commit := Commit{
			CommitMessage: cm,
		}
		commits = append(commits, commit)
	}

	// for _, repoCommit := range repoCommits {
	// 	fmt.Printf("%q\t%q\n", repoCommit.Hash.String()[:7], repoCommit.Message)
	// 	filesIter, err := repoCommit.Files()
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// 	filesIter.ForEach(func(file *object.File) error {
	// 		fmt.Printf("%q\n", file.Name)
	// 		return nil
	// 	})
	// }
	fmt.Println("===")

	return &Version{}, nil
}

// func (s *Service) BuildChangelog(from, to string) (*Changelog, error) {
// 	cl := Changelog{}
// 	return &cl, nil
// }

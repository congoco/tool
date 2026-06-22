package congoco

import (
	"fmt"
	"strings"

	"github.com/go-git/go-git/v6/plumbing/object"
)

type CongocoRepository interface {
	GetVersion() (string, error)
	GetCommits() ([]*object.Commit, error)
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
			invalidCommit := fmt.Sprintf("%s>>%s", c.Hash.String()[:7], invalidMessage)
			invalidCommits = append(invalidCommits, invalidCommit)
		}
	}
	if !valid {
		return invalidCommits, fmt.Errorf("Invalid commits in branch")
	}
	return nil, nil
}

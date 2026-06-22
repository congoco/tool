package congoco

import (
	"fmt"
	"strings"
)

type CongocoRepository interface {
	GetVersion() (string, error)
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

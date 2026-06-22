package congoco

import (
	"fmt"
	"strings"
)

type Service struct{}

func NewService() *Service {
	s := Service{}
	return &s
}

func (s *Service) ParseMessage(message string) (*CommitMessage, error) {
	cm := CommitMessage{}
	subject, _, _ := strings.Cut(message, "\n")
	subject = strings.TrimSpace(subject)
	cm.Subject = subject

	header, _, found := strings.Cut(subject, ":")
	if !found {
		return nil, fmt.Errorf("Invalid commit message")
	}

	header, _, found = strings.Cut(header, "!")
	cm.BreakingChange = found

	cTypeStr, scopeStr, found := strings.Cut(header, "(")
	cType := CommitType(cTypeStr)
	if !found {
		cm.Type = cType
		return &cm, nil
	}

	scope, _, found := strings.Cut(scopeStr, ")")
	if !found {
		return nil, fmt.Errorf("Invalid scope")
	}

	cm.Scope = scope

	return &cm, nil
}

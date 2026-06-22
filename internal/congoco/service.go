package congoco

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

// func (s *Service) ParseMessage(message string) (*CommitMessage, error) {
// 	cm := CommitMessage{}
// 	subject, _, _ := strings.Cut(message, "\n")
// 	subject = strings.TrimSpace(subject)
// 	cm.Subject = subject

// 	header, _, found := strings.Cut(subject, ":")
// 	if !found {
// 		return nil, fmt.Errorf("Invalid commit message")
// 	}

// 	header, _, found = strings.Cut(header, "!")
// 	cm.BreakingChange = found

// 	cTypeStr, scopeStr, found := strings.Cut(header, "(")
// 	cType := CommitType(cTypeStr)
// 	if !found {
// 		cm.Type = cType
// 		return &cm, nil
// 	}

// 	scope, _, found := strings.Cut(scopeStr, ")")
// 	if !found {
// 		return nil, fmt.Errorf("Invalid scope")
// 	}

// 	cm.Scope = scope

// 	return &cm, nil
// }

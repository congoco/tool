package config

const rootPckgName string = "_pckg"

type Service struct {
	repository ConfigRepository
}

type ConfigRepository interface {
	GetDefaults(*Parameters) (*Parameters, error)
	GetVersion() (string, error)
}

func NewService() *Service {
	configRepository := NewRepository()
	s := Service{
		repository: configRepository,
	}
	return &s
}

func (s *Service) LoadParameters() (*Parameters, error) {
	params := NewParameters()
	params, err := s.repository.GetDefaults(params)
	if err != nil {
		return nil, err
	}

	return params, nil
}

func (s *Service) LoadVersion() (string, error) {
	version, err := s.repository.GetVersion()
	if err != nil {
		return "", err
	}
	return version, nil
}

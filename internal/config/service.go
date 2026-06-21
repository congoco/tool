package config

const rootPckgName string = "_pckg"

type Service struct {
	repository ConfigRepository
}

type ConfigRepository interface {
	GetDefaults(*Parameters) (*Parameters, error)
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

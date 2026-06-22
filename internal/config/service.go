package config

const rootPckgName string = "_pckg"

type ConfigRepository interface {
	GetDefaults(params *Config) (*Config, error)
	GetCustomYaml(cfg *Config, customYamlFile bool) (*Config, error)
}

type Service struct {
	repository ConfigRepository
}

func NewService() *Service {
	configRepository := NewRepository()
	s := Service{
		repository: configRepository,
	}
	return &s
}

func (s *Service) LoadDefaults(cfg *Config) (*Config, error) {
	cfg, err := s.repository.GetDefaults(cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}

func (s *Service) LoadCustom(cfg *Config, customYamlFile bool) (*Config, error) {
	cfg, err := s.repository.GetCustomYaml(cfg, customYamlFile)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}

// func (s *Service) LoadVersion() (string, error) {
// 	version, err := s.repository.GetVersion()
// 	if err != nil {
// 		return "", err
// 	}
// 	return version, nil
// }

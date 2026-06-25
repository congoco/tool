package config

type ConfigRepository interface {
	GetDefaults(params *Config) (*Config, error)
	GetCustomYaml(cfg *Config, customYamlFile bool) (*Config, error)
	SaveConfig(cfg *Config, force bool) error
}

type Service struct {
	repo ConfigRepository
}

func NewService() *Service {
	configRepository := NewRepository()
	s := Service{
		repo: configRepository,
	}
	return &s
}

func (s *Service) LoadDefaults(cfg *Config) (*Config, error) {
	cfg, err := s.repo.GetDefaults(cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}

func (s *Service) LoadCustom(cfg *Config, customYamlFile bool) (*Config, error) {
	cfg, err := s.repo.GetCustomYaml(cfg, customYamlFile)
	if err != nil {
		return nil, err
	}

	if cfg.Packages == nil {
		cfg.Packages = make(map[string]Package)
	}

	if cfg.RootPackageEnabled {
		cfg.Packages[cfg.RootPackageName] = Package{}
	}

	for _, pckg := range cfg.Packages {
		if pckg.ChangelogFileName == "" {
			pckg.ChangelogFileName = cfg.ChangelogFilename
		}

		if pckg.ChangelogPath == "" {
			pckg.ChangelogPath = pckg.Path
		}

		pckg.Include = append(pckg.Include, pckg.Path)
	}

	return cfg, nil
}

func (s *Service) CreateConfigFile(configFilename string, force bool) error {
	cfg := NewConfig()
	var err error
	cfg, err = s.LoadDefaults(cfg)
	if err != nil {
		return err
	}
	cfg.CustomConfigFilename = configFilename
	err = s.repo.SaveConfig(cfg, force)
	if err != nil {
		return err
	}
	return nil
}

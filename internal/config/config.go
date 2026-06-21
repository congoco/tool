package config

const versionJsonFileName string = "version.json"

type Config struct {
	Service ConfigService
	*Parameters
	Version string
}

type ConfigService interface {
	LoadDefaults() (*Parameters, error)
	LoadCustomParameters(params *Parameters) (*Parameters, error)
	LoadVersion() (string, error)
}

func New() (*Config, error) {
	configService := NewService()

	version, err := configService.LoadVersion()
	if err != nil {
		return nil, err
	}

	params, err := configService.LoadDefaults()
	if err != nil {
		return nil, err
	}

	cfg := Config{
		Service:    configService,
		Parameters: params,
		Version:    version,
	}

	return &cfg, nil
}

func (c *Config) Reload() error {
	var err error
	c.Parameters, err = c.Service.LoadCustomParameters(c.Parameters)
	if err != nil {
		return err
	}
	return nil
}

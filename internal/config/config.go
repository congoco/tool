package config

const versionJsonFileName string = "version.json"

type Config struct {
	service ConfigService
	*Parameters
	Version string
}

type ConfigService interface {
	LoadParameters() (*Parameters, error)
	LoadVersion() (string, error)
}

func New() (*Config, error) {
	configService := NewService()
	params, err := configService.LoadParameters()
	if err != nil {
		return nil, err
	}
	version, err := configService.LoadVersion()
	if err != nil {
		return nil, err
	}

	cfg := Config{
		service:    configService,
		Parameters: params,
		Version:    version,
	}

	return &cfg, nil
}

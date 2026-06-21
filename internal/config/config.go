package config

type Config struct {
	service ConfigService
	*Parameters
}

type ConfigService interface {
	LoadParameters() (*Parameters, error)
}

func New() (*Config, error) {
	configService := NewService()
	params, err := configService.LoadParameters()
	if err != nil {
		panic(err)
	}

	c := Config{
		service:    configService,
		Parameters: params,
	}

	return &c, nil
}

package config

const CustomConfigPath string = "congoco.yaml"

type Parameters struct {
	ChangelogFilename  string `yaml:"changelog_filename"`
	CustomConfigPath   string `yaml:"-"`
	Formatter          string `yaml:"formatter"`
	LogLevel           string `yaml:"log_level"`
	RootPackageEnabled bool   `yaml:"root_package_enabled"`
	TagPrefix          string `yaml:"tag_prefix"`
}

func NewParameters() *Parameters {
	params := Parameters{
		CustomConfigPath: CustomConfigPath,
	}
	return &params
}

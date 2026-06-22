package config

const (
	versionJsonFilename  string = "version.json"
	customConfigFilename string = "congoco.yaml"
)

type Config struct {
	ChangelogFilename    string `yaml:"changelog_filename"`
	CustomConfigFilename string `yaml:"-"`
	Formatter            string `yaml:"formatter"`
	LogLevel             string `yaml:"log_level"`
	RootPackageEnabled   bool   `yaml:"root_package_enabled"`
	TagPrefix            string `yaml:"tag_prefix"`
	VersionJsonFilename  string `yaml:"-"`
}

func NewConfig() *Config {
	c := Config{
		CustomConfigFilename: customConfigFilename,
		VersionJsonFilename:  versionJsonFilename,
	}
	return &c
}

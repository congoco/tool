package config

type Parameters struct {
	LogLevel           string `yaml:"log_level"`
	Formatter          string `yaml:"formatter"`
	ChangelogFilename  string `yaml:"changelog_filename"`
	RootPackageEnabled bool   `yaml:"root_package_enabled"`
	TagPrefix          string `yaml:"tag_prefix"`
}

func NewParameters() *Parameters {
	params := Parameters{}
	return &params
}

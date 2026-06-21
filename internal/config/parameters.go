package config

type Parameters struct {
	RootPackageEnabled bool   `yaml:"root_package_enabled"`
	TagPrefix          string `yaml:"tag_prefix"`
	ChangelogFilename  string `yaml:"changelog_filename"`
}

func NewParameters() *Parameters {
	params := Parameters{}
	return &params
}

package config

const (
	versionJsonFilename  string = "version.json"
	customConfigFilename string = "congoco.yaml"
	rootPackageName      string = "_root"
)

type Package struct {
	ChangelogFileName string   `yaml:"changelog_filename"`
	ChangelogPath     string   `yaml:"changelog_path"`
	Include           []string `yaml:"include"`
	Path              string   `yaml:"path"`
}

type Config struct {
	ChangelogFilename    string             `yaml:"changelog_filename"`
	CustomConfigFilename string             `yaml:"-"`
	Formatter            string             `yaml:"formatter"`
	Packages             map[string]Package `yaml:"packages"`
	RootPackageEnabled   bool               `yaml:"root_package_enabled"`
	RootPackageName      string             `yaml:"-"`
	TagPrefix            string             `yaml:"tag_prefix"`
	VersionJsonFilename  string             `yaml:"-"`
}

func NewConfig() *Config {
	c := Config{
		CustomConfigFilename: customConfigFilename,
		RootPackageName:      rootPackageName,
		VersionJsonFilename:  versionJsonFilename,
	}
	return &c
}

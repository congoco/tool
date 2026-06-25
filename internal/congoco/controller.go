package congoco

import (
	"os"

	"congoco/internal/config"
	"congoco/internal/flags"

	"github.com/spf13/cobra"
)

type CongocoService interface {
	LoadVersion() (string, error)
	ParseMessage(message string) (*CommitMessage, error)
	ValidateBranch() ([]string, error)
	GetPackageVersions() (map[string]*Version, error)
	CalculatePackagesVersions(invalidCommitsStrategy, belongCommitsStrategy string) (map[string]*Version, error)
	// BuildChangelog(from, to string) (*Changelog, error)
}

type ConfigService interface {
	LoadDefaults(cfg *config.Config) (*config.Config, error)
	LoadCustom(cfg *config.Config, customYamlFile bool) (*config.Config, error)
	CreateConfigFile(configFilename string, force bool) error
}

type Controller struct {
	cfg        *config.Config
	cfgService ConfigService
	flags      flags.Flags
	service    CongocoService
	View       *View
	RootCmd    *cobra.Command
}

func NewController() (*Controller, error) {
	cfg := config.NewConfig()
	flags := flags.Flags{}
	service, err := NewService(cfg)
	if err != nil {
		return nil, err
	}

	c := Controller{
		cfg:        cfg,
		cfgService: nil,
		flags:      flags,
		service:    service,
		RootCmd:    nil,
		View:       nil,
	}

	err = c.bootstrap()
	if err != nil {
		return nil, err
	}
	return &c, nil
}

func (c *Controller) bootstrap() error {
	rootCmd := &cobra.Command{
		Use:               "congoco",
		Short:             "Conventional commits version manager",
		Long:              "Tool for calculating and managing versions from conventional commits.",
		SilenceUsage:      true,
		SilenceErrors:     true,
		PersistentPreRunE: c.preRun,
		Run:               c.root,
	}
	var formatterType flags.FormatterType = flags.TXT
	rootCmd.PersistentFlags().StringVarP(&c.flags.Persistent.Config, "config", "c", c.cfg.CustomConfigFilename, "path to config file")
	rootCmd.PersistentFlags().VarP(&formatterType, "formatter", "o", "output formatter [ini, json, txt]")
	c.RootCmd = rootCmd

	versionCmd := &cobra.Command{
		Use:   "version",
		Short: "Show congoco version",
		Long:  "Show congoco version",
		Run:   c.version,
	}
	c.RootCmd.AddCommand(versionCmd)

	initCmd := &cobra.Command{
		Use:   "init",
		Short: "Create user config file",
		Long:  "Create user config file",
		Run:   c.init,
	}
	initCmd.Flags().BoolVarP(&c.flags.Init.Force, "overwrite", "w", false, "overwrite an existing file")
	c.RootCmd.AddCommand(initCmd)

	validateCmd := &cobra.Command{
		Use:   "validate",
		Short: "Validate Conventional Commits",
		Long:  "Validate Conventional Commits",
		Run:   c.validate,
	}
	validateCmd.Flags().StringVarP(&c.flags.Validate.Message, "message", "m", "", "validate commit message (use single quotes)")
	c.RootCmd.AddCommand(validateCmd)

	currentCmd := &cobra.Command{
		Use:   "current",
		Short: "Show current package versions",
		Long:  "Scan the branch for version tags of all packages",
		Run:   c.current,
	}
	c.RootCmd.AddCommand(currentCmd)

	nextCmd := &cobra.Command{
		Use:   "next",
		Short: "Calculate next package versions",
		Long:  "Scan commits until previous version tag and calculate next Semantic Version",
		Run:   c.next,
	}
	var belongStrategy flags.BelongStrategy = flags.ALL
	nextCmd.Flags().VarP(&belongStrategy, "belong", "b", "package matching strategy [all, scope, path]")

	var invalidStrategy flags.InvalidCommitsStrategy = flags.FAIL
	nextCmd.Flags().VarP(&invalidStrategy, "invalid", "i", "invalid commits handling strategy [fail, ignore, other]")

	nextCmd.Flags().BoolVarP(&c.flags.Next.Changelog, "changelog", "l", false, "create changelog")
	nextCmd.Flags().BoolVarP(&c.flags.Next.VersionFileUpdate, "file-update", "u", false, "update version file")
	nextCmd.Flags().BoolVarP(&c.flags.Next.Commit, "commit", "m", false, "commit changes")
	nextCmd.Flags().BoolVarP(&c.flags.Next.Push, "push", "p", false, "push changes (auto --commit)")
	c.RootCmd.AddCommand(nextCmd)

	// changelogCmd := &cobra.Command{
	// 	Use:   "changelog",
	// 	Short: "Show package changelogs",
	// 	Long:  "Show package changelogs",
	// 	Run:   c.changelog,
	// }
	// changelogCmd.Flags().StringVarP(&c.flags.Changelog.From, "from", "f", "HEAD", "start changelog from [HEAD, tag name, commit hash]")
	// changelogCmd.Flags().StringVarP(&c.flags.Changelog.From, "to", "t", "", "finish changelog on [tag name, commit hash, INIT] (default \"last version tag\")")
	// changelogCmd.Flags().StringVarP(&c.flags.Changelog.Invalid, "invalid", "i", "fail", "invalid commits handling strategy [fail, ignore, other]")
	// c.RootCmd.AddCommand(changelogCmd)

	return nil
}

func (c *Controller) preRun(cmd *cobra.Command, args []string) error {
	if cmd.Name() == "completion" {
		return nil
	}

	c.cfgService = config.NewService()

	var err error
	c.cfg, err = c.cfgService.LoadDefaults(c.cfg)
	if err != nil {
		return err
	}

	yamlFilenameOverwrited := cmd.Flag("config").Changed
	if yamlFilenameOverwrited {
		c.cfg.CustomConfigFilename = c.flags.Persistent.Config
	}

	c.cfg, err = c.cfgService.LoadCustom(c.cfg, yamlFilenameOverwrited)
	if err != nil {
		return err
	}

	if cmd.Flag("formatter").Changed {
		c.cfg.Formatter = c.flags.Persistent.Formatter
	}

	c.View, err = NewView(flags.FormatterType(c.cfg.Formatter))
	if err != nil {
		return err
	}

	return nil
}

func (c *Controller) root(cmd *cobra.Command, args []string) {
	cmd.Help()
}

func (c *Controller) version(cmd *cobra.Command, args []string) {
	output := Output{}
	version, err := c.service.LoadVersion()
	if err != nil {
		output["Error"] = err.Error()
		c.View.Show(output)
		os.Exit(2)
	}
	output["Version"] = version
	c.View.Show(output)
}

func (c *Controller) init(cmd *cobra.Command, args []string) {
	output := Output{}
	force := c.flags.Init.Force
	err := c.cfgService.CreateConfigFile(c.cfg.CustomConfigFilename, force)
	if err != nil {
		output["Error"] = err.Error()
		c.View.Show(output)
		os.Exit(2)
	}
}

func (c *Controller) validate(cmd *cobra.Command, args []string) {
	output := Output{}
	if cmd.Flag("message").Changed {
		if len(c.flags.Validate.Message) < 1 {
			output["Error"] = "Empty message"
			c.View.Show(output)
			os.Exit(2)
		}
		commitMessage, err := c.service.ParseMessage(c.flags.Validate.Message)
		if err != nil {
			output["Error"] = err.Error()
			c.View.Show(output)
			os.Exit(2)
		}
		output["commit"] = commitMessage
		c.View.Show(output)
		return
	}

	invalidCommits, err := c.service.ValidateBranch()
	if err != nil {
		output["Error"] = err.Error()
		output["Invalid commits"] = invalidCommits
		c.View.Show(output)
		os.Exit(2)
	}
}

func (c *Controller) current(cmd *cobra.Command, args []string) {
	output := Output{}
	versions, err := c.service.GetPackageVersions()
	if err != nil {
		output["Error"] = err.Error()
		c.View.Show(output)
	}

	packages := make(map[string]map[string]string, len(versions))

	for pckg, version := range versions {
		tag := ""
		if version.String() != "0.0.0" {
			tag = version.Tag.Name
		}
		packages[pckg] = map[string]string{
			"Version": version.String(),
			"Tag":     tag,
		}
	}

	output["packages"] = packages
	c.View.Show(output)
}

func (c *Controller) next(cmd *cobra.Command, args []string) {
	output := Output{}

	versions, err := c.service.CalculatePackagesVersions(c.flags.Next.InvalidStrategy, c.flags.Next.BelongsStrategy)
	if err != nil {
		output["Error"] = err.Error()
		c.View.Show(output)
	}

	packages := make(map[string]map[string]string, len(versions))

	for pckgName, version := range versions {
		tag := ""

		if version.String() != "0.0.0" {
			tag = version.Tag.Name
		}

		packages[pckgName] = map[string]string{
			"Version": version.String(),
			"Tag":     tag,
		}
	}

	output["packages"] = packages
	c.View.Show(output)
}

// func (c *Controller) changelog(cmd *cobra.Command, args []string) {
// 	output := Output{}
// 	_, err := c.service.CalculatePackagesVersions(c.cfg)
// 	if err != nil {
// 		output["Error"] = err.Error()
// 		c.View.Show(output)
// 	}
// }

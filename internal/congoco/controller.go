package congoco

import (
	"os"

	"congoco/internal/config"

	"github.com/spf13/cobra"
)

type CongocoService interface {
	LoadVersion() (string, error)
}

type ConfigService interface {
	LoadDefaults(cfg *config.Config) (*config.Config, error)
	LoadCustom(cfg *config.Config, customYamlFile bool) (*config.Config, error)
	CreateConfigFile(configFilename string, force bool) error
}

type Controller struct {
	cfg        *config.Config
	cfgService ConfigService
	flags      Flags
	service    CongocoService
	View       *View
	RootCmd    *cobra.Command
}

func NewController() (*Controller, error) {
	cfg := config.NewConfig()
	flags := Flags{}
	service, err := NewService()
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
	rootCmd.PersistentFlags().StringVarP(&c.flags.Persistent.Config, "config", "c", c.cfg.CustomConfigFilename, "path to config file")
	rootCmd.PersistentFlags().StringVarP(&c.flags.Persistent.Formatter, "formatter", "f", string(TXT), "output formatter")

	c.RootCmd = rootCmd

	// == // == //

	versionCmd := &cobra.Command{
		Use:   "version",
		Short: "Show congoco version",
		Long:  "Show congoco version",
		Run:   c.version,
	}
	c.RootCmd.AddCommand(versionCmd)

	// == // == //

	initCmd := &cobra.Command{
		Use:   "init",
		Short: "Create custom config file",
		Long:  "Create custom config file",
		Run:   c.init,
	}
	initCmd.Flags().BoolVarP(&c.flags.Init.Force, "overwrite", "w", false, "overwrite an existing file")
	c.RootCmd.AddCommand(initCmd)

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

	c.View, err = NewView(ViewType(c.cfg.Formatter))
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

package congoco

import (
	"congoco/internal/config"

	"github.com/spf13/cobra"
)

type CongocoService interface {
	LoadVersion() (string, error)
}

type Controller struct {
	cfg     *config.Config
	flags   Flags
	service CongocoService
	view    *View
	RootCmd *cobra.Command
}

func NewController() (*Controller, error) {
	cfg := config.NewConfig()
	flags := Flags{}
	service, err := NewService()
	if err != nil {
		return nil, err
	}

	c := Controller{
		cfg:     cfg,
		flags:   flags,
		service: service,
		view:    nil,
		RootCmd: nil,
	}

	err = c.init()
	if err != nil {
		return nil, err
	}
	return &c, nil
}

func (c *Controller) init() error {
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
		RunE:  c.version,
	}
	c.RootCmd.AddCommand(versionCmd)

	return nil
}

func (c *Controller) preRun(cmd *cobra.Command, args []string) error {
	if cmd.Name() == "completion" {
		return nil
	}

	cfgService := config.NewService()

	var err error
	c.cfg, err = cfgService.LoadDefaults(c.cfg)
	if err != nil {
		return err
	}

	yamlFilenameOverwrited := cmd.Flag("config").Changed
	if yamlFilenameOverwrited {
		c.cfg.CustomConfigFilename = c.flags.Persistent.Config
	}

	c.cfg, err = cfgService.LoadCustom(c.cfg, yamlFilenameOverwrited)
	if err != nil {
		return err
	}

	if cmd.Flag("formatter").Changed {
		c.cfg.Formatter = c.flags.Persistent.Formatter
	}

	c.view, err = NewView(ViewType(c.cfg.Formatter))
	if err != nil {
		return err
	}

	return nil
}

func (c *Controller) root(cmd *cobra.Command, args []string) {
	cmd.Help()
}

func (c *Controller) version(cmd *cobra.Command, args []string) error {
	version, err := c.service.LoadVersion()
	if err != nil {
		return err
	}
	output := Output{
		"Version": version,
	}
	c.view.Show(output)
	return nil
}

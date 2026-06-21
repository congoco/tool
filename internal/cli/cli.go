package cli

import (
	"log/slog"

	"congoco/internal/config"
	"congoco/internal/format"

	"github.com/spf13/cobra"
)

type Cli struct {
	cfg     *config.Config
	log     *slog.Logger
	service CliService
	Flags   Flags
	RootCmd *cobra.Command
}

type CliService interface {
	PreRun(cmd *cobra.Command, flags Persistent) error
	Root(cmd *cobra.Command)
	Version()
	Init(params *config.Parameters, force bool) error
	Validate()
	Current()
	Next()
}

func New(defaultCfg *config.Config) *Cli {
	cliService := NewService(defaultCfg)

	cli := Cli{
		// log:     logger,
		cfg:     defaultCfg,
		service: cliService,
	}

	rootCmd := &cobra.Command{
		Use:   "congoco [flags] [command]",
		Short: "Conventional commits version manager",
		Long:  "Tool for calculating and managing versions from conventional commits.",
		Run: func(cmd *cobra.Command, args []string) {
			cli.service.Root(cmd)
		},
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			err := cli.service.PreRun(cmd, cli.Flags.Persistent)
			if err != nil {
				panic(err)
			}
			return nil
		},
	}

	rootCmd.PersistentFlags().StringVarP(&cli.Flags.Persistent.Config, "config", "c", config.CustomConfigPath, "path to config file")
	rootCmd.PersistentFlags().StringVarP(&cli.Flags.Persistent.Formatter, "formatter", "f", string(format.TXT), "output formatter")

	cli.RootCmd = rootCmd

	cli.init()

	return &cli
}

func (c *Cli) init() {
	versionCmd := &cobra.Command{
		Use:   "version",
		Short: "Show Congoco version",
		Long:  "Version long",
		Run: func(cmd *cobra.Command, args []string) {
			c.service.Version()
		},
	}
	c.RootCmd.AddCommand(versionCmd)

	initCmd := &cobra.Command{
		Use:   "init",
		Short: "Create custom config YAML file",
		Long:  "Init long",
		RunE: func(cmd *cobra.Command, args []string) error {
			params, err := c.cfg.Service.LoadDefaults()
			if err != nil {
				return err
			}
			err = c.service.Init(params, c.Flags.Init.Force)
			if err != nil {
				return err
			}
			return nil
		},
	}
	initCmd.Flags().BoolVarP(&c.Flags.Init.Force, "overwrite", "w", false, "overwrite existed config")
	c.RootCmd.AddCommand(initCmd)

	validateCmd := &cobra.Command{
		Use:   "validate",
		Short: "Validate Conventional Commits in repository",
		Long:  "Check the commit history for compliance with the Conventional Commits specification.",
		Run: func(cmd *cobra.Command, args []string) {
			c.service.Validate()
		},
	}
	c.RootCmd.AddCommand(validateCmd)

	currentCmd := &cobra.Command{
		Use:   "current",
		Short: "Show current version in repository",
		Long:  "Current long",
		Run: func(cmd *cobra.Command, args []string) {
			c.service.Current()
		},
	}
	c.RootCmd.AddCommand(currentCmd)

	nextCmd := &cobra.Command{
		Use:   "next",
		Short: "Calculate next version",
		Long:  "Next long",
		Run: func(cmd *cobra.Command, args []string) {
			c.service.Next()
		},
	}
	c.RootCmd.AddCommand(nextCmd)
}

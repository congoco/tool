package cli

import (
	"log/slog"

	"congoco/internal/config"
	"congoco/internal/format"

	"github.com/spf13/cobra"
)

type Cli struct {
	log     *slog.Logger
	RootCmd *cobra.Command
	service CliService
}

type CliService interface {
	PreRun(cmd *cobra.Command, args []string) error
	Root(cmd *cobra.Command, args []string)
	Version(cmd *cobra.Command, args []string)
	Validate(cmd *cobra.Command, args []string)
	Current(cmd *cobra.Command, args []string)
	Next(cmd *cobra.Command, args []string)
}

func New(defaultCfg *config.Config) *Cli {
	cliService := NewService(defaultCfg)

	cli := Cli{
		// log:     logger,
		service: cliService,
	}

	rootCmd := &cobra.Command{
		Use:               "congoco [flags] [command]",
		Short:             "Conventional commits version manager",
		Long:              "Tool for calculating and managing versions from conventional commits.",
		Run:               cli.service.Root,
		PersistentPreRunE: cli.service.PreRun,
	}

	rootCmd.PersistentFlags().StringVarP(&cliService.Flags.Persistent.Config, "config", "c", config.CustomConfigPath, "path to config file")
	rootCmd.PersistentFlags().StringVarP(&cliService.Flags.Persistent.Formatter, "formatter", "f", string(format.TXT), "output formatter")

	cli.RootCmd = rootCmd

	cli.init()

	return &cli
}

func (c *Cli) init() {
	versionCmd := &cobra.Command{
		Use:   "version",
		Short: "Show Congoco version",
		Long:  "CVersion long",
		Run:   c.service.Version,
	}
	c.RootCmd.AddCommand(versionCmd)

	validateCmd := &cobra.Command{
		Use:   "validate",
		Short: "Validate Conventional Commits in repository",
		Long:  "Check the commit history for compliance with the Conventional Commits specification.",
		Run:   c.service.Validate,
	}
	c.RootCmd.AddCommand(validateCmd)

	currentCmd := &cobra.Command{
		Use:   "current",
		Short: "Show current version in repository",
		Long:  "Current long",
		Run:   c.service.Current,
	}
	c.RootCmd.AddCommand(currentCmd)

	nextCmd := &cobra.Command{
		Use:   "next",
		Short: "Calculate next version",
		Long:  "Next long",
		Run:   c.service.Next,
	}
	c.RootCmd.AddCommand(nextCmd)
}

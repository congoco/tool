package cli

import (
	"log/slog"

	"congoco/internal/config"

	"github.com/spf13/cobra"
)

type Cli struct {
	log     *slog.Logger
	service CliService
	RootCmd *cobra.Command
}

type CliService interface {
	Root(cmd *cobra.Command, args []string)
	Validate(cmd *cobra.Command, args []string)
	Current(cmd *cobra.Command, args []string)
	Next(cmd *cobra.Command, args []string)
}

func New(cfg *config.Config, logger *slog.Logger) *Cli {
	cliService := NewService(cfg)

	cli := Cli{
		log:     logger,
		service: cliService,
	}

	rootCmd := &cobra.Command{
		Use:   "congoco",
		Short: "Conventional commits version manager",
		Long:  "Tool for calculating and managing versions from conventional commits.",
		Run:   cli.service.Root,
	}

	rootCmd.Flags().BoolVarP(&cliService.Flags.Root.Version, "version", "v", false, "congoco version")

	cli.RootCmd = rootCmd

	cli.init()

	return &cli
}

func (c *Cli) init() {
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

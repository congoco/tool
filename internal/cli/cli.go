package cli

import (
	"log/slog"

	"github.com/spf13/cobra"
)

type Cli struct {
	log     *slog.Logger
	service CliService
	RootCmd *cobra.Command
}

type CliService interface {
	Current(cmd *cobra.Command, args []string)
	Next(cmd *cobra.Command, args []string)
}

func New(logger *slog.Logger) *Cli {
	rootCmd := &cobra.Command{
		Use:   "congoco",
		Short: "Conventional commits version manager",
		Long:  "Tool for calculating and managing versions from conventional commits.",
	}

	rootCmd.Flags().BoolP("version", "v", false, "congoco tool version")

	cliService := NewService()

	cli := Cli{
		log:     logger,
		RootCmd: rootCmd,
		service: cliService,
	}

	cli.init()

	return &cli
}

func (c *Cli) init() {
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

package cli

import (
	"github.com/spf13/cobra"
)

type Cli struct {
	Service CliService
	RootCmd *cobra.Command
}

type CliService interface {
	Current(cmd *cobra.Command, args []string)
	Next(cmd *cobra.Command, args []string)
}

func New() *Cli {
	rootCmd := &cobra.Command{
		Use:   "congoco",
		Short: "Conventional commits version manager",
		Long:  "Tool for calculating and managing versions from conventional commits.",
	}

	rootCmd.Flags().BoolP("version", "v", false, "congoco tool version")

	cliService := NewService()

	cli := Cli{
		RootCmd: rootCmd,
		Service: cliService,
	}

	cli.init()

	return &cli
}

func (c *Cli) init() {
	currentCmd := &cobra.Command{
		Use:   "current",
		Short: "Show current version in repository",
		Long:  "Current long",
		Run:   c.Service.Current,
	}
	c.RootCmd.AddCommand(currentCmd)

	nextCmd := &cobra.Command{
		Use:   "next",
		Short: "Calculate next version",
		Long:  "Next long",
		Run:   c.Service.Next,
	}
	c.RootCmd.AddCommand(nextCmd)
}

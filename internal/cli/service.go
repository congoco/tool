package cli

import (
	"congoco/internal/config"
	"congoco/internal/format"

	"github.com/spf13/cobra"
)

type Service struct {
	formatter *format.Formatter
	cfg       *config.Config
	Flags     Flags
}

func NewService(cfg *config.Config, formatter *format.Formatter) *Service {
	s := Service{
		formatter: formatter,
		cfg:       cfg,
	}
	return &s
}

func (s *Service) Root(cmd *cobra.Command, args []string) {
	if s.Flags.Root.Version {
		output := format.Output{
			"Version": s.cfg.Version,
		}
		s.formatter.Render(output)
		return
	}
	cmd.Help()
}

func (s *Service) Validate(cmd *cobra.Command, args []string) {
	panic("<cli.Service.Validate> not implemented")
}

func (s *Service) Current(cmd *cobra.Command, args []string) {
	panic("<cli.Service.Current> not implemented")
}

func (s *Service) Next(cmd *cobra.Command, args []string) {
	panic("<cli.Service.Next> not implemented")
}

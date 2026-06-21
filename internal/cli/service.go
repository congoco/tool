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

func NewService(defaultCfg *config.Config) *Service {
	formatter, err := format.New(defaultCfg.Formatter)
	if err != nil {
		panic(err)
	}
	s := Service{
		formatter: formatter,
		cfg:       defaultCfg,
	}
	return &s
}

func (s *Service) PreRun(cmd *cobra.Command, args []string) error {
	if cmd.Name() == "completion" {
		return nil
	}

	if cmd.Flag("config").Changed {
		s.cfg.CustomConfigPath = s.Flags.Persistent.Config
	}

	err := s.cfg.Reload()
	if err != nil {
		panic(err)
	}

	if cmd.Flag("formatter").Changed {
		s.cfg.Formatter = s.Flags.Persistent.Formatter
	}
	formatter, err := format.New(s.cfg.Formatter)
	if err != nil {
		panic(err)
	}
	s.formatter = formatter

	return nil
}

func (s *Service) Root(cmd *cobra.Command, args []string) {
	cmd.Help()
}

func (s *Service) Version(cmd *cobra.Command, args []string) {
	output := format.Output{
		"Version": s.cfg.Version,
	}
	s.formatter.Render(output)
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

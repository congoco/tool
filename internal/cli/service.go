package cli

import (
	"fmt"

	"congoco/internal/config"
	"congoco/internal/format"

	"github.com/spf13/cobra"
)

type CliRepository interface {
	SaveConfig(*config.Parameters, bool) error
}

type Service struct {
	cfg       *config.Config
	formatter *format.Formatter
	repo      CliRepository
}

func NewService(defaultCfg *config.Config) *Service {
	repo, err := NewRepository()
	if err != nil {
		panic(err)
	}

	formatter, err := format.New(defaultCfg.Formatter)
	if err != nil {
		panic(err)
	}
	s := Service{
		cfg:       defaultCfg,
		formatter: formatter,
		repo:      repo,
	}
	return &s
}

func (s *Service) PreRun(cmd *cobra.Command, flags Persistent) error {
	if cmd.Name() == "completion" {
		return nil
	}

	if cmd.Flag("config").Changed {
		s.cfg.CustomConfigPath = flags.Config
	}

	err := s.cfg.Reload()
	if err != nil {
		panic(err)
	}

	if cmd.Flag("formatter").Changed {
		s.cfg.Formatter = flags.Formatter
	}
	fmt.Printf("Formatter: %s\n", s.cfg.Formatter)
	formatter, err := format.New(s.cfg.Formatter)
	if err != nil {
		panic(err)
	}
	s.formatter = formatter

	return nil
}

func (s *Service) Root(cmd *cobra.Command) {
	cmd.Help()
}

func (s *Service) Version() {
	output := format.Output{
		"Version": s.cfg.Version,
	}
	s.formatter.Render(output)
}

func (s *Service) Init(params *config.Parameters, force bool) error {
	err := s.repo.SaveConfig(params, force)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) Validate() {
	panic("<cli.Service.Validate> not implemented")
}

func (s *Service) Current() {
	panic("<cli.Service.Current> not implemented")
}

func (s *Service) Next() {
	panic("<cli.Service.Next> not implemented")
}

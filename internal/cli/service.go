package cli

import (
	"fmt"

	"congoco/internal/config"

	"github.com/spf13/cobra"
)

type Service struct {
	cfg   *config.Config
	Flags Flags
}

func NewService(cfg *config.Config) *Service {
	s := Service{
		cfg: cfg,
	}
	return &s
}

func (s *Service) Root(cmd *cobra.Command, args []string) {
	if s.Flags.Root.Version {
		fmt.Println(s.cfg.Version)
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

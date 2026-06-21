package cli

import (
	"github.com/spf13/cobra"
)

type Service struct{}

func NewService() *Service {
	s := Service{}
	return &s
}

func (s *Service) Current(cmd *cobra.Command, args []string) {
	panic("<cli.Service.Current> not implemented")
}

func (s *Service) Next(cmd *cobra.Command, args []string) {
	panic("<cli.Service.Next> not implemented")
}

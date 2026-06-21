package congoco

import (
	"fmt"

	"github.com/spf13/cobra"
)

type Service struct{}

func NewService() *Service {
	service := Service{}
	return &service
}

func (s *Service) Hello(cmd *cobra.Command, args []string) {
	fmt.Println("Hello")
}

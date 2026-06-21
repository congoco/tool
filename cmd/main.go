package main

import (
	"fmt"

	"congoco/internal/cli"
	"congoco/internal/config"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		panic(err)
	}

	fmt.Println(cfg.Parameters)

	commands := cli.New()
	err = commands.RootCmd.Execute()
	if err != nil {
		panic(err)
	}
}

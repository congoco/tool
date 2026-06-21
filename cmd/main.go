package main

import (
	"congoco/internal/cli"
	"congoco/internal/config"
	"congoco/internal/logger"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		panic(err)
	}

	log, err := logger.New(cfg.LogLevel)
	if err != nil {
		panic(err)
	}

	cliLogger := log.With("package", "cli")
	commands := cli.New(cfg, cliLogger)
	err = commands.RootCmd.Execute()
	if err != nil {
		panic(err)
	}
}

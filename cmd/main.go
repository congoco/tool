package main

import (
	"os"

	"congoco/internal/cli"
	"congoco/internal/config"
)

func main() {
	defaultCfg, err := config.New() // Full config load in internal/cli/service.go > PreRun
	if err != nil {
		panic(err)
	}

	// log, err := logger.New(cfg.LogLevel)
	// if err != nil {
	// 	panic(err)
	// }

	// cliLogger := log.With("package", "cli")
	commands := cli.New(defaultCfg)
	err = commands.RootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

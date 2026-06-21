package main

import (
	"congoco/internal/cli"
	"congoco/internal/config"
)

func main() {
	defaultCfg, err := config.New()
	if err != nil {
		panic(err)
	}

	// formatter, err := format.New(cfg.Formatter)
	// if err != nil {
	// 	panic(err)
	// }

	// log, err := logger.New(cfg.LogLevel)
	// if err != nil {
	// 	panic(err)
	// }

	// cliLogger := log.With("package", "cli")
	commands := cli.New(defaultCfg)
	err = commands.RootCmd.Execute()
	if err != nil {
		panic(err)
	}
}

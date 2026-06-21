package main

import "congoco/internal/cli"

func main() {
	commands := cli.New()
	err := commands.RootCmd.Execute()
	if err != nil {
		panic(err)
	}
}

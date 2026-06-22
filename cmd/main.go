package main

import (
	"fmt"
	"os"

	"congoco/internal/congoco"
)

func main() {
	cgcController, err := congoco.NewController()
	if err != nil {
		panic(err)
	}
	err = cgcController.RootCmd.Execute()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(2)
	}
}

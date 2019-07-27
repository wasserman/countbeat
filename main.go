package main

import (
	"os"

	"github.com/wasserman/countbeat/cmd"

	_ "github.com/wasserman/countbeat/include"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

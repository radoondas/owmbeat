package main

import (
	"os"

	"github.com/radoondas/owmbeat/cmd"

	_ "github.com/radoondas/owmbeat/include"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

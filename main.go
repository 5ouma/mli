package main

import (
	"os"

	"github.com/5ouma/mli/cmd"
)

func main() {
	cmd := cmd.New()
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}

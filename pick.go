package main

import (
	"os"

	"github.com/ysv/pick/cmd"
)


var (
	version = "0.0.1"
)

func main() {
	err := cmd.Run(version)
	if err != nil {
		os.Exit(1)
	}

	os.Exit(0)
}

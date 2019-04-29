package main

import (
	"os"

	"github.com/ysv/pick/pkg/cli"
)


var (
	version = "0.0.1"
)

func main() {
	err := cli.Run(version)
	if err != nil {
		os.Exit(1)
	}

	os.Exit(0)
}

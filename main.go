package main

import (
	"os"

	"github.com/doucol/whisker-cli/cmd"
)

func main() {
	os.Exit(cmd.Execute())
}

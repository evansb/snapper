package main

import (
	"os"

	"github.com/evansb/snapper/cmd"
)

func main() {
	app := cmd.Snapper()
	app.Run(os.Args)
}

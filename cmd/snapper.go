package cmd

import "gopkg.in/urfave/cli.v1"

const appName = "snapper"

// Snapper returns the CLI App instance of the "snapper" command
func Snapper() *cli.App {
	app := cli.NewApp()
	app.Name = appName

	return app
}

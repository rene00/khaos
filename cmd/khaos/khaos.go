package main

import (
	"github.com/rene00/khaos/internal/commands"
	"github.com/urfave/cli"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "khaos"
	app.Flags = commands.GlobalFlags
	app.Commands = []cli.Command{
		commands.StartCommand,
	}
	app.Run(os.Args)
}

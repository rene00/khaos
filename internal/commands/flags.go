package commands

import "github.com/urfave/cli"

var GlobalFlags = []cli.Flag{
	cli.BoolFlag{
		Name:   "debug",
		Usage:  "run in debug mode",
		EnvVar: "KHAOS_DEBUG",
	},
	cli.StringFlag{
		Name:   "config-file, c",
		Usage:  "load configuration from `FILENAME`",
		Value:  "/etc/khaos/khaos.yml",
		EnvVar: "KHAOS_CONFIG_FILE",
	},
}

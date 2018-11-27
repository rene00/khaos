package commands

import (
	"fmt"
	"github.com/rene00/khaos/internal/khaos"
	"github.com/rene00/khaos/internal/server"
	"github.com/urfave/cli"
	"log"
)

var StartCommand = cli.Command{
	Name:   "start",
	Usage:  "Starts API server",
	Flags:  startFlags,
	Action: startAction,
}

var startFlags = []cli.Flag{
	cli.IntFlag{
		Name:   "server-port, p",
		Usage:  "HTTP server port",
		Value:  8000,
		EnvVar: "KHAOS_SERVER_PORT",
	},
	cli.StringFlag{
		Name:   "server-host, i",
		Usage:  "HTTP server host",
		Value:  "127.0.0.1",
		EnvVar: "KHAOS_SERVER_HOST",
	},
	cli.StringFlag{
		Name:   "server-mode, m",
		Usage:  "debug, release or test",
		Value:  "debug",
		EnvVar: "KHAOS_SERVER_MODE",
	},
	cli.IntFlag{
		Name:   "read-timeout",
		Usage:  "HTTP Read Timeout",
		Value:  60,
		EnvVar: "KHAOS_HTTP_READ_TIMEOUT",
	},
	cli.IntFlag{
		Name:   "write-timeout",
		Usage:  "HTTP Write Timeout",
		Value:  60,
		EnvVar: "KHAOS_HTTP_WRITE_TIMEOUT",
	},
	cli.StringFlag{
		Name:   "database-type",
		Usage:  "Database Type",
		Value:  "sqlite3",
		EnvVar: "KHAOS_DATABASE_TYPE",
	},
	cli.StringFlag{
		Name:   "database-uri",
		Usage:  "Database URI",
		Value:  "./khaos.db",
		EnvVar: "KHAOS_DATABASE_URI",
	},
	cli.StringFlag{
		Name:   "admin-username",
		Usage:  "Admin username",
		Value:  "khaos",
		EnvVar: "KHAOS_ADMIN_USERNAME",
	},
	cli.StringFlag{
		Name:   "admin-password",
		Usage:  "Admin password",
		Value:  "khaos",
		EnvVar: "KHAOS_ADMIN_PASSWORD",
	},
}

func startAction(context *cli.Context) error {
	conf := khaos.NewConfig(context)

	if context.IsSet("server-host") || conf.ServerIP == "" {
		conf.ServerIP = context.String("server-host")
	}

	if context.IsSet("server-port") || conf.ServerPort == 0 {
		conf.ServerPort = context.Int("server-port")
	}

	if context.IsSet("server-mode") || conf.ServerMode == "" {
		conf.ServerMode = context.String("server-mode")
	}

	if conf.ServerPort < 1 {
		log.Fatal("Server port must be a positive integer")
	}

	if context.IsSet("database-type") || conf.DatabaseType == "" {
		conf.DatabaseType = context.String("database-type")
	}

	if context.IsSet("database-uri") || conf.DatabaseURI == "" {
		conf.DatabaseURI = context.String("database-uri")
	}

	if context.IsSet("admin-username") || conf.AdminUsername == "" {
		conf.AdminUsername = context.String("admin-username")
	}

	if context.IsSet("admin-password") || conf.AdminPassword == "" {
		conf.AdminPassword = context.String("admin-password")
	}

	fmt.Printf("Starting API server at %s:%d...\n", context.String("server-host"), context.Int("server-port"))

	server.Start(conf)

	fmt.Println("Done.")
	return nil
}

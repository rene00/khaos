package khaos

import (
	"github.com/urfave/cli"
)

type Config struct {
	Debug         bool
	ConfigFile    string
	ServerIP      string
	ServerPort    int
	ServerMode    string
	DatabaseType  string
	DatabaseURI   string
	AdminUsername string
	AdminPassword string
}

func NewConfig(context *cli.Context) *Config {
	c := &Config{}
	c.SetValuesFromCliContext(context)
	return c
}

func (c *Config) SetValuesFromCliContext(context *cli.Context) error {
	if context.GlobalBool("debug") {
		c.Debug = context.GlobalBool("debug")
	}
	return nil
}

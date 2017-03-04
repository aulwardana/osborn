package main

import (
	"flag"
	"osborn/core/config"
)

var (
	DevMode    bool
	configPath string
)

func initConfig() error {
	flag.BoolVar(&DevMode, "d", true, "Development Mode")
	flag.StringVar(&configPath, "c", "config.yaml", "Configuration File")
	flag.Parse()

	c, err := config.New(configPath)
	if err != nil {
		return err
	}
	cnf = c

	return err
}

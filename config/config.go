package config

import (
	cfg "github.com/micro/go-config"
	"github.com/micro/go-config/source/env"
	"github.com/micro/go-config/source/file"
	"log"
	"strings"
)


// initConfig tries to load and map the ServiceConfiguration struct
// the sources are sequentially loaded: environment-variables, config-file
func init() {
	Init()
}

// Init loads the configuration from file and then from environment variables
func Init() {
	if err := cfg.Load(
		file.NewSource(),
		env.NewSource(),
	); err != nil {
		if strings.Contains(err.Error(), "no such file") {
			log.Print("missing config.json, use environment variables")
		} else {
			log.Fatal(err.Error())
		}
	}
}


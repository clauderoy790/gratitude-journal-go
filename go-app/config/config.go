package config

import (
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"gopkg.in/yaml.v3"
)

// Will be set by LDFLAGS
var VERSION = "UNSET"
var COMMITHASH = "UNSET"

type Config struct {
	Database struct {
		Name     string `yaml:"name"`
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
	} `yaml:"database"`
	App struct {
		MinPasswordLength int    `yaml:"minPasswordLength"`
		Port              int    `yaml:"port"`
		User              string `yaml:"user"`
		Password          string `yaml:"password"`
	} `yaml:"app"`
	Messages struct {
		NoEntryFound string `yaml:"noEntryFound"`
	} `yaml:"messages"`
}

func Get() Config {
	_, filename, _, _ := runtime.Caller(0)
	f, err := os.Open(strings.ReplaceAll(filename, filepath.Ext(filename), ".yaml"))
	if err != nil {
		log.Fatalln("failed to open config file: " + err.Error())
	}
	defer f.Close()

	var cfg Config
	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&cfg)
	if err != nil {
		panic(err)
	}

	return cfg
}

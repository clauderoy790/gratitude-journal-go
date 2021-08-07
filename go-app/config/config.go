package config

import (
	"gopkg.in/yaml.v3"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

type Config struct {
	App struct {
		MinPasswordLength int `yaml:"minPasswordLength"`
		Port              int `yaml:"port"`
	} `yaml:"app"`
	Database struct {
		Name     string `yaml:"name"`
		Cluster  string `yaml:"cluster"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
	} `yaml:"database"`
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

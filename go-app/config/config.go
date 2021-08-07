package config

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
	"path/filepath"
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
	UseLocalhost bool
}

func Get() Config {
	wd, err := os.Getwd()
	fmt.Println("wd: ", wd)
	if err != nil {
		panic("fail to find wd: " + err.Error())
	}

	configPath := filepath.Join(wd, "config/config.yaml")

	f, err := os.Open(configPath)
	if err != nil {
		panic("failed to open config file: " + err.Error())
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

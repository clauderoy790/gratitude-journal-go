package config

import (
	"gopkg.in/yaml.v3"
	"os"
	"path/filepath"
)

type Config struct {
	App struct {
		MinPasswordLength int `yaml:"minPasswordLength"`
		Port int `yaml:"port"`
	} `yaml:"app"`
	Database struct {
		Name string `yaml:"name"`
		Cluster string `yaml:"cluster"`
		User string `yaml:"user"`
		Password string `yaml:"password"`
	} `yaml:"database"`
	Messages struct {
		NoEntryFound string `yaml:"noEntryFound"`
	} `yaml:"messages"`
}

func Get() Config {
	projetDir,err := filepath.Abs("./")
	if err != nil {
		panic(err)
	}
	f, err := os.Open(projetDir+"/config/config.yaml")
	if err != nil {
		panic(err)
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
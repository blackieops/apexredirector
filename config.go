package main

import (
	"os"

	"gopkg.in/yaml.v2"
)

type RedirectConfig struct {
	FromHost string `yaml:"from_host"`
	ToHost   string `yaml:"to_host"`
}

type Config struct {
	Secure    bool             `yaml:"secure"`
	Redirects []RedirectConfig `yaml:"redirects"`
}

func ReadConfig(configPath string) (*Config, error) {
	config := &Config{}

	file, err := os.Open(configPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	d := yaml.NewDecoder(file)
	if err := d.Decode(&config); err != nil {
		return nil, err
	}

	return config, nil
}

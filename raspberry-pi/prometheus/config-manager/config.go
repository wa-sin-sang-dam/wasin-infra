package main

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Global globalConfig   `yaml:"global"`
	Scrape []scrapeConfig `yaml:"scrape_configs"`
}

type globalConfig struct {
	ScrapeInterval     string `yaml:"scrape_interval"`
	EvaluationInterval string `yaml:"evaluation_interval"`
}

type scrapeConfig struct {
	JobName       string         `yaml:"job_name"`
	StaticConfigs []staticConfig `yaml:"static_configs"`
}

type staticConfig struct {
	Targets []string `yaml:"targets"`
}

func NewConfig(filePath string) (Config, error) {
	config := Config{}

	yamlFile, err := os.ReadFile(*configPath)
	if err != nil {
		return config, err
	}

	err = config.load(yamlFile)
	return config, err
}

func (c *Config) load(bytes []byte) error {
	return yaml.Unmarshal(bytes, c)
}

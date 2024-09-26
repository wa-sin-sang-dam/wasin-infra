package main

import (
	"errors"
	"os"

	"gopkg.in/yaml.v3"
)

var (
	errJobNameNotFound = errors.New("could not find the matching static config: no such job name")
)

type Config struct {
	Global   globalConfig   `yaml:"global"`
	Scrapes  []scrapeConfig `yaml:"scrape_configs"`
	filePath string
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
	config := Config{
		filePath: filePath,
	}

	yamlFile, err := os.ReadFile(*configPath)
	if err != nil {
		return config, err
	}

	err = config.load(yamlFile)
	return config, err
}

func (c *Config) AddTarget(jobName string, ip string) error {
	index, err := c.findScrapeConfigIndex(jobName)
	if err != nil {
		return err
	}
	targets := c.Scrapes[index].StaticConfigs[0].Targets
	targets = append(targets, ip)

	c.Scrapes[index].StaticConfigs[0].Targets = targets
	return nil
}

func (c *Config) ListTargets(jobName string) ([]string, error) {
	index, err := c.findScrapeConfigIndex(jobName)
	if err != nil {
		return nil, err
	}
	return c.Scrapes[index].StaticConfigs[0].Targets, nil
}

func (c *Config) RemoveTarget(jobName string, ip string) error {
	index, err := c.findScrapeConfigIndex(jobName)
	if err != nil {
		return err
	}

	targets := c.Scrapes[index].StaticConfigs[0].Targets
	targets = RemoveItemInSlice(targets, ip)
	c.Scrapes[index].StaticConfigs[0].Targets = targets

	return nil
}

func (c *Config) Save() error {
	bytes, err := yaml.Marshal(c)
	if err != nil {
		return err
	}

	err = os.WriteFile(c.filePath, bytes, 0666)
	if err != nil {
		return err
	}

	return nil
}

func (c *Config) load(bytes []byte) error {
	return yaml.Unmarshal(bytes, c)
}

func (c *Config) findScrapeConfigIndex(jobName string) (int, error) {
	for index, scrapeConfig := range c.Scrapes {
		if scrapeConfig.JobName == jobName {
			return index, nil
		}
	}

	return -1, errJobNameNotFound
}

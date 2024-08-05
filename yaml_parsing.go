package main

import (
	"gopkg.in/yaml.v3"
	"io"
)

type SqlcConfig struct {
	Version string `yaml:"version"`
	Sql     []struct {
		Engine  string `yaml:"engine"`
		Queries string `yaml:"queries"`
		Schema  string `yaml:"schema"`
		Gen     struct {
			Go struct {
				Out string `yaml:"out"`
			} `yaml:"go"`
		} `yaml:"gen"`
	} `yaml:"sql"`
}

func parseSqlcYaml(f io.ReadCloser) (folderPath string, err error) {
	var cfg SqlcConfig
	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&cfg)
	if err != nil {
		return "", err
	}
	return cfg.Sql[0].Gen.Go.Out, nil
}

type CubeConfig struct {
	Version string `yaml:"version"`
	Go      struct {
		Source string `yaml:"source"`
		Target string `yaml:"target"`
	} `yaml:"go"`
}

func parseSqlCube(f io.ReadCloser) (src, target string, err error) {
	var cfg CubeConfig
	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&cfg)
	if err != nil {
		return "", "", err
	}
	return cfg.Go.Source, cfg.Go.Target, nil
}

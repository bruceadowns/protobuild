package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/BurntSushi/toml"
)

const configVersion = "unstable"

type config struct {
	Version   string
	Generator string
	Plugins   []string
	Includes  struct {
		Before   []string
		Vendored []string
		After    []string
	}

	Packages map[string]string

	Descriptors []struct {
		Prefix      string
		Target      string
		IgnoreFiles []string `toml:"ignore_files"`
	}
}

func newDefaultConfig() config {
	return config{
		Generator: "go",
		Includes: struct {
			Before   []string
			Vendored []string
			After    []string
		}{
			Before: []string{"."},
			After:  []string{"/usr/local/include"},
		},
	}
}

func readConfig(path string) (config, error) {
	p, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatalln(err)
	}
	c := newDefaultConfig()
	if err := toml.Unmarshal(p, &c); err != nil {
		log.Fatalln(err)
	}

	if c.Version != configVersion {
		return config{}, fmt.Errorf("unknown file version %v; please upgrade to %v", c.Version, configVersion)
	}

	return c, nil
}

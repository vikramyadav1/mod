package config

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"

	"gopkg.in/yaml.v2"
)

type Modconfig struct {
	D struct {
		Defaults struct {
			Volumes []string `yaml:"volumes"`
			Command string   `yaml:"command"`
			Ports   []string `yaml:"ports"`
		}
	}
}

func NewModConfig() Modconfig {
	homedir, _ := os.UserHomeDir()
	m := Modconfig{}
	data, err := ioutil.ReadFile(path.Join(homedir, ".modconfig"))

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	err = yaml.Unmarshal([]byte(data), &m)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	return m
}

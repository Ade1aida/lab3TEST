package main

import (
	"fmt"
	"os"

	"github.com/Ade1aida/contains/pkg/contains"
	"github.com/go-yaml/yaml"
)

type Config struct {
	Filename  string `yaml:"filename"`
	Substring string `yaml:"substring"`
}

func main() {

	var config Config
	yamlFile, err := os.ReadFile("test.yaml")
	if err != nil {
		panic(err)
	}

	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		panic(err)
	}

	textFile := config.Filename
	substring := config.Substring
	result, err := contains.Contain(textFile, substring)

	if err != nil {
		fmt.Print(err)
	} else {
		print(result)
	}
}

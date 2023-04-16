package cfg

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

type input struct {
	FileName string `yaml:"file_name"`
	OwnURL   string `yaml:"own_url"`
}

func GetInput() *input {
	c := &input{}
	yamlFile, err := ioutil.ReadFile("./input.yaml")
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	return c
}

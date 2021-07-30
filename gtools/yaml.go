package gtools

import (
	"io/ioutil"

	gyaml "gopkg.in/yaml.v2"
)

var Yaml yamlInterface

type yamlInterface interface {
	Load(file string, configStruct interface{}) (interface{}, error)
}

type yaml struct{}

func init() {
	Yaml = &yaml{}
}

func (_yaml yaml) Load(file string, configStruct interface{}) (interface{}, error) {
	yamlFile, err := ioutil.ReadFile(file)
	if err != nil {
		return configStruct, err
	}
	err = gyaml.Unmarshal(yamlFile, configStruct)
	if err != nil {
		return configStruct, err
	}
	return configStruct, nil
}

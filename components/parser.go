package components

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

func ParseYamlConfFromBytes(yamlOriginData []byte) (config Config, err error) {

	if len(yamlOriginData) == 0 {
		log.Fatalln("错误的`config.yml`配置")
	} else {
		err = yaml.Unmarshal(yamlOriginData, &config)
	}

	return
}

func ParseYamlFromFile(yamlFileUri string) (config Config, err error) {

	var fileData []byte

	fileData, err = ioutil.ReadFile(yamlFileUri)

	if err == nil {
		config, err = ParseYamlConfFromBytes(fileData)
	}

	return
}

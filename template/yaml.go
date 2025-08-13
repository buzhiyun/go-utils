package template

import (
	"github.com/buzhiyun/go-utils/log"
	"gopkg.in/yaml.v3"
)

func GetTemplateFromYaml(yamlByte []byte, tepl string) string {
	var data interface{}
	err := yaml.Unmarshal(yamlByte, &data)
	if err != nil {
		log.Errorf("yaml.Unmarshal error: %v", err)
		return ""
	}
	return getTemplate(&data, &tepl)
}

func GetTemplateFromYamlString(yamlstr string, tepl string) string {
	return GetTemplateFromYaml([]byte(yamlstr), tepl)
}

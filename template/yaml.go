package template

import (
	"os"
	"text/template"

	"github.com/buzhiyun/go-utils/log"
	"gopkg.in/yaml.v3"
)

func GetTemplateStrFromYaml(yamlByte []byte, tepl string) string {
	var data interface{}
	err := yaml.Unmarshal(yamlByte, &data)
	if err != nil {
		log.Errorf("yaml.Unmarshal error: %v", err)
		return ""
	}
	return getTemplate(&data, &tepl)
}

func GetTemplateStrFromYamlString(yamlstr string, tepl string) string {
	return GetTemplateStrFromYaml([]byte(yamlstr), tepl)
}

func GetTemplateFromYaml(yamlByte []byte, tepl *template.Template) string {
	var data interface{}
	err := yaml.Unmarshal(yamlByte, &data)
	if err != nil {
		log.Errorf("yaml.Unmarshal error: %v", err)
		return ""
	}
	return getTemplateString(&data, tepl)
}

func GetTemplateFromYamlFile(filename string, tepl *template.Template) string {
	yamlByte, err := os.ReadFile(filename)
	if err != nil {
		log.Errorf("os.ReadFile error: %v", err)
		return ""
	}
	return GetTemplateFromYaml(yamlByte, tepl)
}

func GetTemplateFromYamlString(yamlstr string, tepl *template.Template) string {
	return GetTemplateFromYaml([]byte(yamlstr), tepl)
}

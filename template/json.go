package template

import (
	"text/template"

	"github.com/buzhiyun/go-utils/log"
	jsoniter "github.com/json-iterator/go"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

func GetTemplateStrFromJsonString(jsonstr string, tepl string) string {
	var data interface{}
	err := json.UnmarshalFromString(jsonstr, &data)
	if err != nil {
		log.Errorf("json.UnmarshalFromString error: %v", err)
		return ""
	}

	return getTemplate(&data, &tepl)

}

func GetTemplateStrFromJson(jsonByte []byte, tepl string) string {
	var data interface{}
	err := json.Unmarshal(jsonByte, &data)
	if err != nil {
		log.Errorf("json.Unmarshal error: %v", err)
		return ""
	}
	return getTemplate(&data, &tepl)
}

func GetTemplateFromJsonString(jsonstr string, tepl *template.Template) string {
	var data interface{}
	err := json.UnmarshalFromString(jsonstr, &data)
	if err != nil {
		log.Errorf("json.UnmarshalFromString error: %v", err)
		return ""
	}
	return getTemplateString(&data, tepl)
}

func GetTemplateFromJson(jsonByte []byte, tepl *template.Template) string {
	var data interface{}
	err := json.Unmarshal(jsonByte, &data)
	if err != nil {
		log.Errorf("json.Unmarshal error: %v", err)
		return ""
	}

	return getTemplateString(&data, tepl)
}

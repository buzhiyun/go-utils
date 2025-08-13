package template

import (
	"bytes"
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

func GetTemplateFromJson(jsonByte []byte, tepl string) string {
	var data interface{}
	err := json.Unmarshal(jsonByte, &data)
	if err != nil {
		log.Errorf("json.Unmarshal error: %v", err)
		return ""
	}
	return getTemplate(&data, &tepl)
}

func getTemplate(data *interface{}, tepl *string) string {

	t, err := template.New("test").Parse(*tepl)
	if err != nil {
		log.Errorf("template.New error: %v", err)
		return ""
	}
	var buf bytes.Buffer
	err = t.Execute(&buf, *data)
	if err != nil {
		log.Errorf("t.Execute error: %v", err)
		return ""
	}
	return buf.String()
}

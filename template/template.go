package template

import (
	"bytes"
	"text/template"

	"github.com/buzhiyun/go-utils/log"
)

func getTemplate(data *interface{}, tepl *string) string {

	t, err := template.New("test").Parse(*tepl)
	if err != nil {
		log.Errorf("template.New error: %v", err)
		return ""
	}
	return getTemplateString(data, t)
}

func getTemplateString(data *interface{}, tepl *template.Template) string {
	var buf bytes.Buffer
	err := tepl.Execute(&buf, *data)
	if err != nil {
		log.Errorf("t.Execute error: %v", err)
		return ""
	}
	return buf.String()
}

func GetTemplateString(data *interface{}, tepl *template.Template) string {
	return getTemplateString(data, tepl)
}

package http

import (
	"bytes"
	"github.com/buzhiyun/go-utils/log"
	jsoniter "github.com/json-iterator/go"
	"io"
	"mime/multipart"
	"net/http"
	"strings"
	"time"
)

var http_client = &http.Client{}
var json = jsoniter.ConfigCompatibleWithStandardLibrary

func HttpPostJson(url string, body interface{}) (responseBody []byte, err error) {
	requestJson, err := json.Marshal(body)
	if err != nil {
		log.Errorf("http 发送初始化失败，无法json参数, %v", body)
		return
	}
	//加上协议头
	if !strings.HasPrefix(url, "http://") && !strings.HasPrefix(url, "https://") {
		url = "http://" + url
	}

	log.Debugf("发送接口: %s ，body: %s", url, requestJson)

	req, err := http.NewRequest("POST", url, bytes.NewReader(requestJson))
	if err != nil {
		return
	}

	req.Header.Set("Content-Type", "application/json")
	//req.Header.Set("Cookie", "name=anny")

	http_client.Timeout = 5 * time.Second
	resp, err := http_client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	responseBody, err = io.ReadAll(resp.Body)

	return
}

func HttpGet(url string) (responseBody []byte, err error) {
	//加上协议头
	if !strings.HasPrefix(url, "http://") && !strings.HasPrefix(url, "https://") {
		url = "http://" + url
	}

	resp, err := http.Get(url)
	if err != nil {
		return
	}

	defer resp.Body.Close()
	responseBody, err = io.ReadAll(resp.Body)
	return
}

func HttpPostFile(url string, formField map[string]string, fileName string, fileField string, src io.Reader) (responseBody []byte, err error) {
	method := "POST"

	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)

	//part1, errFile1 := writer.CreateFormField("")
	part1, errFile1 := writer.CreateFormFile(fileField, fileName)
	if errFile1 != nil {
		return nil, errFile1
	}
	io.Copy(part1, src)
	_, errFile1 = io.Copy(part1, src)
	if errFile1 != nil {
		log.Error(errFile1)
		return
	}

	for k, v := range formField {
		_ = writer.WriteField(k, v)
	}
	//_ = writer.WriteField("filename", "filename.txt")
	//_ = writer.WriteField("name", "media")

	err = writer.Close()
	if err != nil {
		log.Errorf("关闭 mime writer 错误: %s", err)
		return
	}

	client := &http.Client{}

	//加上协议头
	if !strings.HasPrefix(url, "http://") && !strings.HasPrefix(url, "https://") {
		url = "http://" + url
	}

	req, err := http.NewRequest(method, url, payload)

	log.Debugf("%s", payload.Bytes())
	if err != nil {
		log.Errorf("创建上传请求错误: %s", err)
		return
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())
	res, err := client.Do(req)
	if err != nil {
		log.Errorf("发送上传请求错误: %s", err)
		return
	}
	defer res.Body.Close()

	responseBody, err = io.ReadAll(res.Body)
	if err != nil {
		log.Errorf("读取上传返回错误: %s", err)
		return
	}
	log.Infof("上传成功，返回： %s", responseBody)
	return
}

func HttpGetWithBasicAuth(url, username, password string) (responseBody []byte, err error) {
	//加上协议头
	if !strings.HasPrefix(url, "http://") && !strings.HasPrefix(url, "https://") {
		url = "http://" + url
	}

	req, err := http.NewRequest("GET", url, nil)

	req.SetBasicAuth(username, password)
	http_client.Timeout = 5 * time.Second
	resp, err := http_client.Do(req)
	if err != nil {
		return
	}

	defer resp.Body.Close()
	responseBody, err = io.ReadAll(resp.Body)
	return
	//fmt.Println(string(body))
}

package http

import (
	"bytes"
	"context"
	"io"
	"mime/multipart"
	"net/http"
	neturl "net/url"

	"github.com/buzhiyun/go-utils/log"
	jsoniter "github.com/json-iterator/go"

	"strings"
	"time"
)

var (
	json        = jsoniter.ConfigCompatibleWithStandardLibrary
	http_client *http.Client
)

func setOption(client *http.Client, option ...HttpClientOption) {
	for _, option := range option {
		if option.Timeout != 0 {
			log.Debugf("http client 设置 timeout %v", option.Timeout)
			client.Timeout = option.Timeout
		}
	}
}

func httpClient(option ...HttpClientOption) *http.Client {
	if http_client != nil {
		setOption(http_client, option...)
		return &http.Client{}
	}
	_client := &http.Client{}
	setOption(_client, option...)
	return _client
}

func HttpPostJson(url string, body interface{}, headers map[string]string, option ...HttpClientOption) (responseBody []byte, err error) {
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
	for k, v := range headers {
		req.Header.Set(k, v)
	}

	resp, err := httpClient(option...).Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	responseBody, err = io.ReadAll(resp.Body)

	return
}

func HttpPostJsonWithCtx(ctx context.Context, url string, body interface{}, headers map[string]string, option ...HttpClientOption) (responseBody []byte, err error) {
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

	req, err := http.NewRequestWithContext(ctx, "POST", url, bytes.NewReader(requestJson))
	if err != nil {
		return
	}

	req.Header.Set("Content-Type", "application/json")
	//req.Header.Set("Cookie", "name=anny")
	for k, v := range headers {
		req.Header.Set(k, v)
	}

	resp, err := httpClient(option...).Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	responseBody, err = io.ReadAll(resp.Body)

	return
}

func HttpPostForm(url string, formData map[string]string, headers map[string]string, option ...HttpClientOption) (responseBody []byte, err error) {
	var body = neturl.Values{}
	for k, v := range formData {
		body.Set(k, v)
	}

	//加上协议头
	if !strings.HasPrefix(url, "http://") && !strings.HasPrefix(url, "https://") {
		url = "http://" + url
	}

	log.Debugf("发送接口: %s ，body: %s", url, body.Encode())
	req, err := http.NewRequest("POST", url, strings.NewReader(body.Encode()))
	if err != nil {
		return
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	for k, v := range headers {
		req.Header.Set(k, v)
	}
	//req.Header.Set("Cookie", "name=anny")

	// http_client.Timeout = 5 * time.Second
	resp, err := httpClient(option...).Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	responseBody, err = io.ReadAll(resp.Body)

	return
}

func HttpPostFormWithCtx(ctx context.Context, url string, formData map[string]string, headers map[string]string, option ...HttpClientOption) (responseBody []byte, err error) {
	var body = neturl.Values{}
	for k, v := range formData {
		body.Set(k, v)
	}

	//加上协议头
	if !strings.HasPrefix(url, "http://") && !strings.HasPrefix(url, "https://") {
		url = "http://" + url
	}

	log.Debugf("发送接口: %s ，body: %s", url, body.Encode())
	req, err := http.NewRequestWithContext(ctx, "POST", url, strings.NewReader(body.Encode()))
	if err != nil {
		return
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	for k, v := range headers {
		req.Header.Set(k, v)
	}
	//req.Header.Set("Cookie", "name=anny")

	// http_client.Timeout = 5 * time.Second
	resp, err := httpClient(option...).Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	responseBody, err = io.ReadAll(resp.Body)

	return
}

func HttpGet(url string, headers map[string]string, option ...HttpClientOption) (responseBody []byte, err error) {
	//加上协议头
	if !strings.HasPrefix(url, "http://") && !strings.HasPrefix(url, "https://") {
		url = "http://" + url
	}

	req, err := http.NewRequest("GET", url, nil)
	log.Debugf("发送接口: GET %s", url)
	if err != nil {
		return
	}
	for k, v := range headers {
		req.Header.Set(k, v)
	}

	res, err := httpClient(option...).Do(req)
	if err != nil {
		return
	}
	defer res.Body.Close()

	responseBody, err = io.ReadAll(res.Body)
	return
}

func HttpGetWithCtx(ctx context.Context, url string, headers map[string]string, option ...HttpClientOption) (responseBody []byte, err error) {
	//加上协议头
	if !strings.HasPrefix(url, "http://") && !strings.HasPrefix(url, "https://") {
		url = "http://" + url
	}

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	log.Debugf("发送接口: GET %s", url)
	if err != nil {
		return
	}
	for k, v := range headers {
		req.Header.Set(k, v)
	}

	res, err := httpClient(option...).Do(req)
	if err != nil {
		return
	}
	defer res.Body.Close()

	responseBody, err = io.ReadAll(res.Body)
	return
}

func HttpPostFile(url string, formField map[string]string, fileName string, fileField string, src io.Reader, headers map[string]string, option ...HttpClientOption) (responseBody []byte, err error) {
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
	for k, v := range headers {
		req.Header.Set(k, v)
	}
	res, err := httpClient(option...).Do(req)
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

func HttpGetWithBasicAuth(url, username, password string, option ...HttpClientOption) (responseBody []byte, err error) {
	//加上协议头
	if !strings.HasPrefix(url, "http://") && !strings.HasPrefix(url, "https://") {
		url = "http://" + url
	}

	req, err := http.NewRequest("GET", url, nil)

	req.SetBasicAuth(username, password)
	http_client.Timeout = 5 * time.Second
	resp, err := httpClient(option...).Do(req)
	if err != nil {
		return
	}

	defer resp.Body.Close()
	responseBody, err = io.ReadAll(resp.Body)
	return
	//fmt.Println(string(body))
}

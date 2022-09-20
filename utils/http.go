package utils

import (
	"bytes"
	"crypto/tls"
	"errors"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

// Get 请求  link：请求url
func Get(link string, header map[string]string, params map[string]string) ([]byte, error) {
	client := &http.Client{Timeout: 5 * time.Second}
	//忽略https的证书
	client.Transport = &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	p := url.Values{}
	u, _ := url.Parse(link)
	// 设置url参数
	for k, v := range params {
		p.Set(k, v)
	}
	u.RawQuery = p.Encode()
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	// 设置header
	for k, v := range header {
		req.Header.Add(k, v)
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("http is exception")
	}
	return io.ReadAll(resp.Body)
}

// PostJSON link：请求url body：json数据
func PostJSON(link string, body string) ([]byte, error) {
	resp, err := http.Post(link, "application/json; charset=utf-8", strings.NewReader(body))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return io.ReadAll(resp.Body)
}

// PostForm link：请求url
func PostForm(link string, values map[string]string) ([]byte, error) {
	p := url.Values{}
	for k, v := range values {
		p.Set(k, v)
	}
	resp, err := http.PostForm(link, p)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return io.ReadAll(resp.Body)
}

// PostFile 上传单个文件 link：请求url
func PostFile(link string, params map[string]string, filename string, path string) ([]byte, error) {
	client := &http.Client{}
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	part, err := writer.CreateFormFile(filename, path)
	if err != nil {
		return nil, err
	}
	_, err = io.Copy(part, file)
	if err != nil {
		return nil, err
	}

	for key, val := range params {
		_ = writer.WriteField(key, val)
	}

	err = writer.Close()
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("POST", link, body)
	if err != nil {
		return nil, err
	}

	request.Header.Set("Content-Type", writer.FormDataContentType())

	resp, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	return io.ReadAll(resp.Body)
}

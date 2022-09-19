package utils

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"time"
)

// Get 发送GET请求
func Get(url string) (string, error) {

	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Set("Accept", "*/*")
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		return "", err
	}

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	result, _ := io.ReadAll(resp.Body)

	return string(result), nil
}

// Post 发送POST请求
func Post(url string, data interface{}, contentType string) ([]byte, error) {
	// 超时时间：10秒
	client := &http.Client{Timeout: 10 * time.Second}
	jsonStr, _ := json.Marshal(data)
	resp, err := client.Post(url, contentType, bytes.NewBuffer(jsonStr))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	result, _ := io.ReadAll(resp.Body)
	return result, nil
}

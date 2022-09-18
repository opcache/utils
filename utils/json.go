package utils

import (
	"bytes"
	"encoding/json"
)

// Decode 解码
func Decode(value string, r interface{}) error {
	buffer := bytes.NewBuffer([]byte(value))
	decoder := json.NewDecoder(buffer)
	return decoder.Decode(r)
}

// Encode 编码
func Encode(value interface{}) (string, error) {
	buffer := bytes.NewBuffer(nil)
	encoder := json.NewEncoder(buffer)
	err := encoder.Encode(value)
	if err != nil {
		return "", err
	}
	return buffer.String(), nil
}

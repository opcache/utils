package utils

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

var chars = []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z",
	"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z",
	"1", "2", "3", "4", "5", "6", "7", "8", "9", "0"}

// 保证每次生成的随机数不一样
func init() {
	rand.Seed(time.Now().UnixNano())
}

// RandString 生成数字和字母
func RandString(lenNum int) string {
	str := strings.Builder{}
	length := len(chars)
	for i := 0; i < lenNum; i++ {
		l := chars[rand.Intn(length)]
		str.WriteString(l)
	}
	return str.String()
}

// RandAlpha 仅生成字符
func RandAlpha(lenNum int) string {
	str := strings.Builder{}
	length := 52
	for i := 0; i < lenNum; i++ {
		l := chars[rand.Intn(length)]
		str.WriteString(l)
	}
	return str.String()
}

// RandNumber 生成指定长度的随机数字
func RandNumber(lenNum int) string {
	str := strings.Builder{}
	length := 10
	numbers := chars[52:62]
	fmt.Println(numbers)
	for i := 0; i < lenNum; i++ {
		str.WriteString(numbers[rand.Intn(length)])
	}
	return str.String()
}

// RandNumberNoZero 生成指定长度的无0随机数字
func RandNumberNoZero(lenNum int) string {
	str := strings.Builder{}
	length := 9
	noZeroNumber := chars[52:61]
	fmt.Println(noZeroNumber)
	for i := 0; i < lenNum; i++ {
		str.WriteString(noZeroNumber[rand.Intn(length)])
	}
	return str.String()
}

// Uuid 生成uuid
func Uuid() (string, error) {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%x-%x-%x-%x-%x", b[0:4], b[4:6], b[6:8], b[8:10], b[10:]), nil
}

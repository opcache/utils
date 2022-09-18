package utils

import (
	"time"
)

//  格式化时间模版

const DATE_TEMPLATE = "2006-01-02"
const DATETIME_TEMPLATE = "2006-01-02 15:04:05"
const DATETIME_UNIXNANO_TEMPLATE = "2006-01-02 15:04:05.000"

// GetNowTime 获取时间戳
func GetNowTime() int64 {
	return time.Now().Local().Unix()
}

// GetMicroTime 获取毫秒时间戳
func GetMicroTime() int64 {
	return time.Now().Local().UnixMilli()
}

// UnixTimeForFormat 时间戳转时间字符串
func UnixTimeForFormat(timeUnix int64, format string) string {
	return time.Unix(timeUnix, 0).Local().Format(format)
}

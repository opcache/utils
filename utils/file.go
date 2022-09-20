package utils

import (
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"path"
)

// GetSize 获取文件大小
func GetSize(f multipart.File) (int, error) {
	content, err := io.ReadAll(f)
	return len(content), err
}

// GetExt 获取文件后缀
func GetExt(fileName string) string {
	return path.Ext(fileName)
}

// CheckExist 检查文件是否存在
func CheckExist(src string) bool {
	_, err := os.Stat(src)
	return os.IsNotExist(err)
}

// CheckPermission 检查文件权限
func CheckPermission(src string) bool {
	_, err := os.Stat(src)

	return os.IsPermission(err)
}

// IsNotExistMkDir 检查文件夹是否存在
// 如果不存在则新建文件夹
func IsNotExistMkDir(src string) error {
	if exist := !CheckExist(src); exist == false {
		if err := MkDir(src); err != nil {
			return err
		}
	}

	return nil
}

// MkDir 新建文件夹
func MkDir(src string) error {
	err := os.MkdirAll(src, os.ModePerm)
	if err != nil {
		return err
	}

	return nil
}

// Open 打开文件
func Open(name string, flag int, perm os.FileMode) (*os.File, error) {
	f, err := os.OpenFile(name, flag, perm)
	defer f.Close()
	if err != nil {
		return nil, err
	}

	return f, nil
}

// Download 下载文件
func Download(remoteUrl, dstFile string) error {
	client := &http.Client{}
	uri, err := url.Parse(remoteUrl)
	if err != nil {
		return err
	}
	// Create the file
	out, err := os.Create(dstFile)
	if err != nil {
		return err
	}
	defer out.Close()

	request, err := http.NewRequest("GET", uri.String(), nil)
	request.Header.Add("Connection", "close")
	request.Header.Add("Host", uri.Host)
	request.Header.Add("Referer", uri.String())
	request.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 6.1; WOW64; rv:12.0) Gecko/20100101 Firefox/12.0")

	resp, err := client.Do(request)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		_, err = io.Copy(out, resp.Body)
	} else {
		return errors.New(fmt.Sprintf("bad status: %s", resp.Status))
	}
	return nil
}

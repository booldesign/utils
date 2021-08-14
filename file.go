package utils

import (
	"os"
	"path/filepath"
)

/**
 * @Author: BoolDesign
 * @Email: booldesign@163.com
 * @Date: 2021/1/27 10:16
 * @Desc:
 */

func OpenFile(filename string) (*os.File, error) {
	f, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0600)
	if err != nil {
		return nil, err
	}
	return f, nil
}

// 判断文件/文件夹是否存在，true存在，false不存在
func FileExists(path string) bool {
	_, err := os.Stat(path) //os.Stat获取文件信息
	if err != nil && !os.IsExist(err) {
		return false
	}
	return true
}

// 判断文件是否存在
func IsFile(path string) bool {
	s, err := os.Stat(path)
	if err != nil { //不是文件也不是目录
		return false
	}
	return !s.IsDir()
}

// 判断是否为文件夹
func IsDir(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		return false
	}
	return s.IsDir()
}

//获取文件大小
func FileSize(file string) (int64, error) {
	f, err := os.Stat(file)
	if err != nil {
		return 0, err
	}
	return f.Size(), nil
}

//获取文件修改时间
func FileMTime(file string) (int64, error) {
	f, err := os.Stat(file)
	if err != nil {
		return 0, err
	}
	return f.ModTime().Unix(), nil
}

// 获取文件扩展名
func FileExt(f string) string {
	ext := filepath.Ext(f)
	if ext == "" {
		return ext
	}
	return ext[1:]
}

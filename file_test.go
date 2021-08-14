package utils

import "testing"

/**
 * @Author: BoolDesign
 * @Email: booldesign@163.com
 * @Date: 2021/8/12 23:03
 * @Desc:
 */

func TestOpenFile(t *testing.T) {
	t.Log(OpenFile("./file.go"))
}

func TestFileExists(t *testing.T) {
	t.Log(FileExists("./file.go"))
}

func TestIsFile(t *testing.T) {
	t.Log(IsFile("./file.go"))
}

package utils

import "testing"

/**
 * @Author: BoolDesign
 * @Email: booldesign@163.com
 * @Date: 2021/2/3 22:30
 * @Desc:
 */

func TestZipFile(t *testing.T) {
	//模拟数据, 也可以读取文件
	var files = []File{
		{"readme.txt", []byte("This archive contains some text files.")},
		{"gopher.txt", []byte("Gopher names:\nGeorge\nGeoffrey\nGonzo")},
		{"todo.txt", []byte("Get animal handling licence.\nWrite more examples.")},
	}

	ZipFile(files, "wei")

}

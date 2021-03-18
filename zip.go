package utils

import (
	"archive/zip"
	"bytes"
	"fmt"
	"log"
	"os"
)

/**
 * @Author: BoolDesign
 * @Email: booldesign@163.com
 * @Date: 2020-05-26 11:57
 * @Desc:
 */

type HeaderSetter interface {
	Header(string, string)
}

//设置文件下载头
func SetZipHeader(headerSetter HeaderSetter, topic string) {
	headerSetter.Header("Content-Type", "application/x-zip-compressed")
	headerSetter.Header("Content-disposition", fmt.Sprintf("attachment; filename=%s", topic))
}

type File struct {
	Name string
	Body []byte
}

//压缩文件
func ZipFile(files []File, filename string) {
	// Create a buffer to write our archive to.
	buf := new(bytes.Buffer)

	// Create a new zip archive.
	w := zip.NewWriter(buf)

	// Add some files to the archive.

	for _, file := range files {
		f, err := w.Create(file.Name)
		if err != nil {
			log.Fatal(err)
		}
		_, err = f.Write([]byte(file.Body))
		if err != nil {
			log.Fatal(err)
		}
	}

	// Make sure to check the error on Close.
	err := w.Close()
	if err != nil {
		log.Fatal(err)
	}

	file, _ := os.Create(filename + ".zip")
	file.WriteString(buf.String())
}

package utils

import (
	"strconv"
	"testing"
)

/**
 * @Author: BoolDesign
 * @Email: booldesign@163.com
 * @Date: 2021/3/19 14:58
 * @Desc:
 */

func TestGenerateAndValidatePassword(t *testing.T) {
	tests := []struct {
		in string
	}{
		{"123123123"},
		{"abc123"},
	}

	for _, test := range tests {
		actual, err := GeneratePassword(test.in)
		if err != nil {
			t.Error(err.Error())
		}

		isOK, err := ValidatePassword(test.in, string(actual))
		if err != nil {
			t.Error(err.Error())
		}

		if !isOK {
			t.Errorf("GeneratePassword and GeneratePassword failed.")
		}
	}
}

func BenchmarkGeneratePassword(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GeneratePassword(strconv.Itoa(i))
	}

	//打印报告
	b.ReportAllocs()
}

func BenchmarkValidatePassword(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ValidatePassword("123123123", "$2a$10$XF4QmSwpUCsjWAtrIect7OKqM2ZYxlIHKVHpUGKouTKozxLeWZr7u")
	}

	//打印报告
	b.ReportAllocs()
}

/*
BenchmarkGeneratePassword-4                           13         141869800 ns/op            5200 B/op         11 allocs/op
BenchmarkValidatePassword-4                            8         153089670 ns/op            5437 B/op         16 allocs/op
*/

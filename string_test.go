package utils

import (
	"testing"
)

/**
 * @Author: BoolDesign
 * @Email: booldesign@163.com
 * @Date: 2021/1/27 13:39
 * @Desc:
 */

func TestToCamel(t *testing.T) {
	tests := []struct {
		x      string
		expect string
	}{
		{"hello world", "HelloWorld"},
		{"hello_world", "HelloWorld"},
		{"hello-world", "HelloWorld"},
	}

	for _, test := range tests {
		if actual := ToCamel(test.x); actual != test.expect {
			t.Errorf("ToCamel(%s): expect %s, actual %s",
				test.x, test.expect, actual)
		}
	}
}

func TestToLowerCamel(t *testing.T) {
	tests := []struct {
		x      string
		expect string
	}{
		{"hello world", "helloWorld"},
		{"hello_world", "helloWorld"},
		{"hello-world", "helloWorld"},
	}

	for _, test := range tests {
		if actual := ToLowerCamel(test.x); actual != test.expect {
			t.Errorf("ToLowerCamel(%s): expect %s, actual %s",
				test.x, test.expect, actual)
		}
	}
}

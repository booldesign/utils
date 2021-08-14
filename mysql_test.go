package utils

import "testing"

/**
 * @Author: BoolDesign
 * @Email: booldesign@163.com
 * @Date: 2021/8/14 12:29
 * @Desc:
 */

func TestDbFieldJoin(t *testing.T) {
	tests := []struct {
		input  []string
		expect string
	}{
		{[]string{"a", "b", "c"}, `'a', 'b', 'c'`},
	}

	for _, test := range tests {
		if actual := DbFieldJoin(test.input); actual != test.expect {
			t.Errorf("DbFieldJoin(%s): expect %s, actual %s",
				test.input, test.expect, actual)
		}
	}
}

func TestFiledDbType2GoType(t *testing.T) {
	tests := []struct {
		input  string
		expect string
	}{
		{"bigint(20)", "int"},
		{"text", "string"},
		{"json", "string"},
		{"tinyint(1)", "int"},
		{"timestamp", "time.Time"},
	}

	for _, test := range tests {
		if actual := FiledDbType2GoType(test.input); actual != test.expect {
			t.Errorf("FiledDbType2GoType(%s): expect %s, actual %s",
				test.input, test.expect, actual)
		}
	}
}

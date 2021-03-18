package utils

import "testing"

/**
 * @Author: BoolDesign
 * @Email: booldesign@163.com
 * @Date: 2021/2/3 22:54
 * @Desc:
 */

func TestWFormatDuration(t *testing.T) {
	t.Log(WFormatDuration(61))
}

func TestConvertToBin(t *testing.T) {
	tests := []struct {
		x      int
		expect string
	}{
		{1, "1"},
		{10, "1010"},
		{25, "11001"},
	}

	for _, test := range tests {
		if actual := ConvertToBin(test.x); actual != test.expect {
			t.Errorf("ConvertToBin(%d): expect %s, actual %s",
				test.x, test.expect, actual)
		}
	}
}

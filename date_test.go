package utils

import (
	"testing"
)

/**
 * @Author: BoolDesign
 * @Email: booldesign@163.com
 * @Date: 2021/1/27 11:58
 * @Desc:
 */

func TestDateToTime(t *testing.T) {
	// 2011-12-31 00:00:00 +0800 CST
	t.Log(DateToTime("2011-12-31"))
}

func TestDateTimeToTime(t *testing.T) {
	// 2011-12-31 06:57:01 +0800 CST
	t.Log(DateTimeToTime("2011-12-31 06:57:01"))
}

func TestTimestampToTime(t *testing.T) {
	// 2021-01-27 11:46:46 +0800 CST
	t.Log(TimestampToTime(1611719206))
}

func TestNowDatetime(t *testing.T) {
	// 2021-03-16 17:15:43
	t.Log(NowDatetime())
}

func TestNowTimestamp(t *testing.T) {
	// 1615886143
	t.Log(NowTimestamp())
}

func TestNowTime(t *testing.T) {
	// 2021-03-16 17:24:56.393992 +0800 CST
	t.Log(NowTime())
}

func TestBirthdayToAge(t *testing.T) {

	tests := []struct {
		x      string
		expect uint8
	}{
		{"2011-12-31", 9}, //TODO：坑
		{"2019-01-01", 2},
	}

	for _, test := range tests {
		if actual := BirthdayToAge(DateToTime(test.x)); actual != test.expect {
			t.Errorf("BirthdayToAge(%s): expect %d, actual %d",
				test.x, test.expect, actual)
		}
	}
}

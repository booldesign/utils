package utils

import (
	"time"
)

/**
 * @Author: BoolDesign
 * @Email: booldesign@163.com
 * @Date: 2021/1/27 10:12
 * @Desc:
 */

const (
	DefaultLocal       = "Asia/Shanghai"
	DefaultDataTime    = "2006-01-02 15:04:05"
	DefaultData        = "2006-01-02"
	DefaultErrData     = "1000-01-01"
	DefaultErrDataTime = "1000-01-01 00:00:00"
)

//go语言并没有全局设置时区这么一个东西，每次输出时间都需要调用一个In()函数改变时区
var loc, _ = time.LoadLocation(DefaultLocal) //设置时区

// 日期年月日 转 Time 零点
func DateToTime(date string) time.Time {
	t, err := time.ParseInLocation(DefaultData, date, loc)
	if err != nil {
		t, _ = time.Parse(DefaultData, DefaultErrData)
	}
	return t
}

// 日期年月日时分秒 转 Time
func DateTimeToTime(datetime string) time.Time {
	t, err := time.ParseInLocation(DefaultDataTime, datetime, loc)
	if err != nil {
		t, _ = time.Parse(DefaultDataTime, DefaultErrDataTime)
	}
	return t
}

// 时间戳 to Time
func TimestampToTime(timestamp int64) time.Time {
	return time.Unix(timestamp, 0).In(loc)
}

// 当前时间
func NowDatetime() string {
	return time.Now().In(loc).Format("2006-01-02 15:04:05")
}

// 当前时间戳
func NowTimestamp() int64 {
	return time.Now().In(loc).Unix()
}

func NowTime() time.Time {
	return time.Now().In(loc)
}

// 生日计算年龄精确到日 周岁
func BirthdayToAge(birthday time.Time) uint8 {
	birthYear, birthMonth, birthDay := birthday.Year(), birthday.Month(), birthday.Day()

	nowTime := time.Now().In(loc)
	year, month, day := nowTime.Year(), nowTime.Month(), nowTime.Day()

	age := year - birthYear
	//满周岁
	if birthMonth > month || (birthMonth == month && birthDay > day) {
		age--
	}

	if age > 128 {
		age = 0
	}

	return uint8(age)
}

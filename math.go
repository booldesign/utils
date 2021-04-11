package utils

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

/**
 * @Author: BoolDesign
 * @Email: booldesign@163.com
 * @Date: 2021/1/27 10:12
 * @Desc:
 */

//范围随机[min, max)
func RandomRangeInt(min, max int) int {
	if min > max {
		max, min = min, max
	}
	if min == max {
		return max
	}
	rand.Seed(time.Now().Unix())
	return rand.Intn(max-min) + min

}

const HOUR = 3600
const MINUTE = 60

//秒 转 时分秒
func WFormatDuration(t int) string {
	str := ""
	last := 0
	if t >= HOUR {
		str += fmt.Sprintf("%d", t/HOUR) + " 时 "
		last = t % HOUR
	} else {
		last = t
	}

	if last >= MINUTE {
		str += fmt.Sprintf("%d", last/MINUTE) + " 分 "
		last = t % MINUTE
	}
	str += fmt.Sprintf("%d", last) + " 秒 "
	return str
}

// 将十进制数字转化为二进制字符串
func ConvertToBin(num int) string {
	s := ""

	if num == 0 {
		return "0"
	}

	// num /= 2 每次循环的时候 都将num除以2  再把结果赋值给 num
	for ; num > 0; num /= 2 {
		lsb := num % 2
		// strconv.Itoa() 将数字强制性转化为字符串
		s = strconv.Itoa(lsb) + s
	}
	return s
}

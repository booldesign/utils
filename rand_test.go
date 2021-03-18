package utils

import "testing"

/**
 * @Author: BoolDesign
 * @Email: booldesign@163.com
 * @Date: 2021/1/30 22:25
 * @Desc:
 */

func TestRandString(t *testing.T) {
	t.Log(RandString(6))
}

func TestRandNumString(t *testing.T) {
	t.Log(RandNumString(6))
}

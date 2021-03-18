package utils

import (
	"testing"
)

/**
 * @Author: BoolDesign
 * @Email: booldesign@163.com
 * @Date: 2021/1/30 18:00
 * @Desc:
 */

func TestGetIntranceIp(t *testing.T) {
	t.Log(GetIntranceIp())
}

func TestIpString2Int(t *testing.T) {
	t.Log(IpString2Int("192.168.199.189"))
}

func TestIpInt2String(t *testing.T) {
	t.Log(IpInt2String(3232286653))
}

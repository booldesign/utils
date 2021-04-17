package utils

import (
	"fmt"
	"strings"
)

/**
 * @Author: BoolDesign
 * @Email: booldesign@163.com
 * @Date: 2021/1/27 15:44
 * @Desc: 动态数据掩码（Dynamic Data Masking，简称为DDM）能够防止把敏感数据暴露给未经授权的用户。
 */

// 手机号 格式：前3后4 示例：135****1436
func MarkMobile(str string) string {
	if len(str) != 11 {
		return str
	}

	return fmt.Sprintf("%s****%s", str[:3], str[len(str)-4:])
}

// 邮箱 格式：前1后1 示例：b**n@gmail.com
func MarkEmail(str string) string {
	if !strings.Contains(str, "@") {
		return str
	}

	split := strings.Split(str, "@")
	if len(split[0]) < 1 || len(split[1]) < 1 {
		return str
	}
	mark := strings.Repeat("*", len(split[0])-2)
	return fmt.Sprintf("%s%s%s@%s", split[0][:1], mark, split[0][len(split[0])-1:], split[1])
}

// 姓名 格式：隐姓 示例：*建文
func MarkRealName(str string) string {
	if len(str) < 1 {
		return ""
	}

	nameRune := []rune(str)
	return fmt.Sprintf("*%s", string(nameRune[1:]))
}

// 密码 ******
func MarkPassWord() string {
	return "******"
}

// 银行卡号 格式：前6后4 示例：621483******3553
func MarkBankCard(str string) string {
	if len(str) > 19 || len(str) < 16 {
		return str
	}

	return fmt.Sprintf("%s******%s", str[:6], str[len(str)-4:])
}

// 身份证号 格式：前10后4 示例：3212811989****0931
func MarkIDCard(str string) string {
	if len(str) != 18 {
		return str
	}

	return fmt.Sprintf("%s****%s", str[:10], str[len(str)-4:])
}

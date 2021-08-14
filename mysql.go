package utils

import "strings"

/**
 * @Author: BoolDesign
 * @Email: booldesign@163.com
 * @Date: 2021/8/13 12:28
 * @Desc:
 */

// DbFieldJoin sql字段加引号和分隔号
func DbFieldJoin(fields []string) string {
	var buf strings.Builder
	for _, str := range fields {
		buf.WriteString("'" + str + "', ")
	}
	return strings.TrimRight(buf.String(), ", ")
}

// FiledDbType2GoType db字段类型转go类型
func FiledDbType2GoType(fieldType string) string {
	fieldType = strings.TrimSpace(fieldType)
	// bigint(20) 转 int
	switch strings.Split(fieldType, "(")[0] {
	case "int", "integer", "mediumint", "bit", "year", "smallint", "tinyint", "bigint":
		return "int"
	case "decimal", "double", "float", "numeric":
		return "float64"
	case "timestamp", "datetime", "time":
		return "time.Time"
	case "json": // TODO: 先转string
		return "string"
	default:
		return "string"
	}
}

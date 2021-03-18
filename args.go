package utils

import "os"

/**
 * @Author: BoolDesign
 * @Email: booldesign@163.com
 * @Date: 2021/1/27 15:44
 * @Desc: 目前仅es里用到
 */

func GetArgsValue(idx int, def string) string {
	if len(os.Args) > idx {
		return os.Args[idx]
	}
	return def
}

//[]int to []interface{}
func IntSliceToArgs(ids []int) []interface{} {
	args := make([]interface{}, len(ids))
	for k, orgId := range ids {
		args[k] = orgId
	}
	return args
}

//[]string to []interface{}
func StringSliceToArgs(strs []string) []interface{} {
	args := make([]interface{}, len(strs))
	for k, s := range strs {
		args[k] = s
	}
	return args
}

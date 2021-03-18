package utils

import (
	"bytes"
	"reflect"
	"sort"
	"strconv"
	"strings"
)

/**
 * @Author: BoolDesign
 * @Email: booldesign@163.com
 * @Date: 2021/1/27 10:36
 * @Desc:
 */

//string slice去重
func StringSliceUnique(a []string) (ret []string) {
	sort.Strings(a)
	va := reflect.ValueOf(a)
	for i := 0; i < va.Len(); i++ {
		if i > 0 && reflect.DeepEqual(va.Index(i-1).String(), va.Index(i).String()) {
			continue
		}
		ret = append(ret, va.Index(i).String())
	}
	return ret
}

//int slice去重
func IntSliceUnique(a []int) (ret []int) {
	sort.Ints(a)
	va := reflect.ValueOf(a)
	for i := 0; i < va.Len(); i++ {
		if i > 0 && reflect.DeepEqual(va.Index(i-1).Int(), va.Index(i).Int()) {
			continue
		}
		ret = append(ret, int(va.Index(i).Int()))
	}
	return ret
}

//整型切片组合为字符串
func IntSliceJoin(a []int, delim string) string {
	var buf bytes.Buffer
	for _, v := range a {
		buf.WriteString(strconv.Itoa(v))
		buf.WriteString(delim)
	}
	return strings.TrimRight(buf.String(), delim)
}

// 比较slice，返回差集，slice1 在 slice2中没有的值
func IntSliceDiff(slice1, slice2 []int) (diffSlice []int) {
	for _, v := range slice1 {
		if !InSliceInt(v, slice2) {
			diffSlice = append(diffSlice, v)
		}
	}
	return
}

// 比较slice，返回交集
func IntSliceIntersect(slice1, slice2 []int) (diffslice []int) {
	for _, v := range slice1 {
		if InSliceInt(v, slice2) {
			diffslice = append(diffslice, v)
		}
	}
	return
}

// v是否包含在sl里
func InSliceInt(v int, sl []int) bool {
	for _, vv := range sl {
		if vv == v {
			return true
		}
	}
	return false
}

// v是否包含在sl里
func InSliceString(v string, sl []string) bool {
	for _, vv := range sl {
		if vv == v {
			return true
		}
	}
	return false
}

//删除Slice指定元素
func RemoveSliceElement(sl []int, v int) []int {
	findIndex := -1
	for index, val := range sl {
		if val == v {
			findIndex = index
			break
		}
	}

	if findIndex != -1 {
		sl = append(sl[0:findIndex], sl[findIndex+1:]...)
	}

	return sl
}

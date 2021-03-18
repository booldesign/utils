package utils

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

/**
 * @Author: BoolDesign
 * @Email: booldesign@163.com
 * @Date: 2021/1/27 11:12
 * @Desc:
 */

// 获取环境变量
func getEnv(key string) string {
	val := os.Getenv(key)
	if val == "" {
		panic(fmt.Sprintf("ENV [%s] not found", key))
	} else if val == "-" {
		val = ""
	}
	return val
}

// 字符串型变量
func GetStringEnv(key string) string {
	val := getEnv(key)
	resVal := val
	return resVal
}

func GetStringEnvDefault(key string, def string) string {
	val := os.Getenv(key)
	if val == "" || val == "-" {
		return def
	}
	return val
}

// 整型变量
func GetIntEnv(key string) int {
	val := getEnv(key)
	resVal, err := strconv.Atoi(val)
	if err != nil {
		panic(err)
	}
	return resVal
}

func GetIntEnvDefault(key string, def int) int {
	val := os.Getenv(key)
	if val == "" || val == "-" {
		return def
	}
	resVal, err := strconv.Atoi(val)
	if err != nil {
		return def
	}
	return resVal
}

// 布尔变量
func GetBoolEnv(key string) bool {
	val := getEnv(key)
	resVal, err := strconv.ParseBool(val)
	if err != nil {
		panic(err)
	}
	return resVal
}

func GetBoolEnvDefault(key string, def bool) bool {
	val := getEnv(key)
	if val == "" || val == "-" {
		return def
	}
	resVal, err := strconv.ParseBool(val)
	if err != nil {
		panic(err)
	}
	return resVal
}

//获取集群地址
func GetStringSliceEnv(key string) []string {
	val := getEnv(key)
	return strings.Split(val, ",")
}

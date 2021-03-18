package utils

import (
	"fmt"

	jsoniter "github.com/json-iterator/go"
)

/**
 * @Author: BoolDesign
 * @Email: booldesign@163.com
 * @Date: 2021/1/27 11:52
 * @Desc:
 */

var json = jsoniter.ConfigCompatibleWithStandardLibrary

func JsonEncode(v interface{}) []byte {
	encoded, err := json.Marshal(v)
	if err != nil {
		return nil
	}
	return encoded
}

func JsonEncodeToString(v interface{}) string {
	encoded, err := json.Marshal(v)
	if err != nil {
		return ""
	}
	return string(encoded)
}

func JsonDecode(data []byte, v interface{}) error {
	err := json.Unmarshal(data, v)
	if err != nil {
		fmt.Errorf("json unmarshal error:%s, data:%s", err.Error(), string(data))
		return err
	}
	return nil
}

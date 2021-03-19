package utils

import (
	"errors"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"strings"
)

/**
 * @Author: BoolDesign
 * @Email: booldesign@163.com
 * @Date: 2021/3/19 14:48
 * @Desc:
 */

//Deprecated
//加密密码(太耗时，逐步废弃)
func GeneratePassword(userPassword string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(userPassword), bcrypt.DefaultCost)
}

//Deprecated
//验证密码(太耗时，逐步废弃)
func ValidatePassword(userPassword string, hashed string) (isOk bool, err error) {
	if err = bcrypt.CompareHashAndPassword([]byte(hashed), []byte(userPassword)); err != nil {
		return false, errors.New("密码比对错误")
	}
	return true, nil
}

//密码加密(新)
func GeneratePasswordNew(password, salt string) string {
	suffix := RandNumString(6)
	hash := Md5(salt + password + suffix)
	return fmt.Sprintf("%s_%s", hash, suffix)
}

// 验证密码(新)
func ValidatePasswordNew(password, hashed, salt string) bool {
	passwords := strings.Split(hashed, "_")
	if len(passwords) != 2 {
		return false
	}
	return Md5(salt+password+passwords[1]) == passwords[0]
}

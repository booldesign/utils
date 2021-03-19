package utils

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
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

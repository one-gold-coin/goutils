package goutils

import (
	"crypto/md5"
	"encoding/hex"
	"reflect"
	"regexp"
)

func Md5(encrypt string) string {
	h := md5.New()
	h.Write([]byte(encrypt))
	return hex.EncodeToString(h.Sum(nil))
}

//是否邮箱
func IsMail(mail string) bool {
	isMatch, err := regexp.MatchString("^([.a-zA-Z0-9_-])+@([a-zA-Z0-9_-])+((.[a-zA-Z0-9_-]{2,10}){1,3})$", mail)
	if err != nil {
		return false
	}
	if isMatch {
		return true
	} else {
		return false
	}
}

//是否手机号，11位数字
func IsMobile(str string) bool {
	isMatch, err := regexp.MatchString("^1[0-9]{10}$", str)
	if err != nil {
		return false
	}
	if isMatch {
		return true
	} else {
		return false
	}
}

// 返回空结构体
func EmptyStruct() struct{} {
	return struct{}{}
}

// 返回空结构体
func EmptySlice() []struct{} {
	return []struct{}{}
}

// IsNil
func IsNil(d interface{}) bool {
	return reflect.ValueOf(d).IsNil()
}

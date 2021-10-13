package goutils

import (
	"strconv"
	"strings"
)

func StrConv(val string) *StringVal {
	return &StringVal{Val: val}
}

type StringVal struct {
	Val string
}

func (m *StringVal) String() string {
	return m.Val
}

func (m *StringVal) StrSlice() []string {
	valList := strings.Split(m.Val, ",")
	return valList
}

func (m *StringVal) Int() int {
	if m.Val == "" {
		return 0
	}
	valInt, _ := strconv.Atoi(m.Val)
	return valInt
}

func (m *StringVal) IntSlice() []int {
	if m.Val == "" {
		return nil
	}
	valList := strings.Split(m.Val, ",")
	valIntList := make([]int, 0)
	for _, s := range valList {
		valInt, _ := strconv.Atoi(s)
		valIntList = append(valIntList, valInt)
	}
	return valIntList
}

func (m *StringVal) Int64() int64 {
	if m.Val == "" {
		return 0
	}
	valInt64, _ := strconv.ParseInt(m.Val, 10, 64)
	return valInt64
}

func (m *StringVal) Int64Slice() []int64 {
	if m.Val == "" {
		return nil
	}
	valList := strings.Split(m.Val, ",")
	valIntList := make([]int64, 0)
	for _, s := range valList {
		valInt64, _ := strconv.ParseInt(s, 10, 64)
		valIntList = append(valIntList, valInt64)
	}
	return valIntList
}

// StringIsEmpty 字符串是否为空
func StringIsEmpty(str ...string) bool {
	for _, s := range str {
		if s != "" {
			return false
		}
	}
	return true
}

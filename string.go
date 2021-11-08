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
	Err error
}

func (m *StringVal) Error() error {
	return m.Err
}

func (m *StringVal) String() string {
	return m.Val
}

func (m *StringVal) StrSlice() []string {
	v := strings.Split(m.Val, ",")
	return v
}

func (m *StringVal) Int() int {
	if m.Val == "" {
		return 0
	}
	v, err := strconv.Atoi(m.Val)
	m.Err = err
	return v
}

func (m *StringVal) IntSlice() []int {
	if m.Val == "" {
		return nil
	}
	vList := strings.Split(m.Val, ",")
	vIList := make([]int, 0)
	for _, s := range vList {
		vI, err := strconv.Atoi(s)
		if err != nil {
			m.Err = err
			return nil
		}
		vIList = append(vIList, vI)
	}
	return vIList
}

func (m *StringVal) Int64() int64 {
	if m.Val == "" {
		return 0
	}
	vI, err := strconv.ParseInt(m.Val, 10, 64)
	m.Err = err
	return vI
}

func (m *StringVal) Int64Slice() []int64 {
	if m.Val == "" {
		return nil
	}
	vList := strings.Split(m.Val, ",")
	vIList := make([]int64, 0)
	for _, s := range vList {
		vI, err := strconv.ParseInt(s, 10, 64)
		if err != nil {
			m.Err = err
			return nil
		}
		vIList = append(vIList, vI)
	}
	return vIList
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

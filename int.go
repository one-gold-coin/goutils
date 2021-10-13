package goutils

import "strconv"

func IntConv(val int) *IntVal {
	return &IntVal{Val: val}
}

type IntVal struct {
	Val int
}

func (m *IntVal) String() string {
	return strconv.Itoa(m.Val)
}

func (m *IntVal) Int() int {
	return m.Val
}

func (m *IntVal) Int64() int64 {
	return int64(m.Val)
}

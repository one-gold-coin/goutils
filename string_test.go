package goutils

import (
	"fmt"
	"testing"
)

func Test_StrConv(t *testing.T) {
	fmt.Println("int: ", StrConv("1").Int())
	fmt.Println("int64: ", StrConv("1").Int64())
	fmt.Println("string: ", StrConv("1").String())
	fmt.Println("StrSlice: ", StrConv("1,2,3").StrSlice())
}

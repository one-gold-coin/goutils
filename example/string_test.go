package main

import (
	"fmt"
	"github.com/one-gold-coin/goutils"
	"testing"
)

func Test_StrConv(t *testing.T) {
	fmt.Println("int: ", goutils.StrConv("1").Int())
	fmt.Println("int64: ", goutils.StrConv("1").Int64())
	fmt.Println("string: ", goutils.StrConv("1").String())
	fmt.Println("StrSlice: ", goutils.StrConv("1,2,3").StrSlice())
	fmt.Println("StringIsEmpty: ", goutils.StringIsEmpty(""))

	//err
	strConv := goutils.StrConv("a")
	fmt.Println("Int64 : ", strConv.Int64())
	fmt.Println("Int64 Err: ", strConv.Error())


}

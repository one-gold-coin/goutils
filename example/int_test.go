package main

import (
	"fmt"
	"github.com/one-gold-coin/goutils"
	"testing"
)

func Test_IntConv(t *testing.T) {
	fmt.Println("int: ", goutils.IntConv(1).Int())
	fmt.Println("int64: ", goutils.IntConv(1).Int64())
	fmt.Println("string: ", goutils.IntConv(1).String())
}

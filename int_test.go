package goutils

import (
	"fmt"
	"testing"
)

func Test_IntConv(t *testing.T) {
	fmt.Println("int: ", IntConv(1).Int())
	fmt.Println("int64: ", IntConv(1).Int64())
	fmt.Println("string: ", IntConv(1).String())
}

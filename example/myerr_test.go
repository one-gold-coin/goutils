package main

import (
	"fmt"
	"github.com/one-gold-coin/goutils"
	"testing"
)

func Test_MyErr(t *testing.T) {
	fmt.Printf("err: %+v \n", goutils.MyErr(1))
	fmt.Printf("err: %+v \n", goutils.MyErr(2, "success"))
	fmt.Printf("err: %+v \n", goutils.MyErr(3, "success", 1))
	fmt.Printf("err: %+v \n", goutils.MyErr(4, "success", struct {
	}{}))
}

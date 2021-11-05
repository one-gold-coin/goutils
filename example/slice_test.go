package main

import (
	"fmt"
	"github.com/one-gold-coin/goutils"
	"testing"
)

func Test_SliceUnique(t *testing.T) {
	sliceList := []int64{0, 1, 1, 2, 2, 3, 4, 5, 1, 23, 4, 56}
	r := goutils.SliceUnique(sliceList)
	fmt.Printf("r: %+v", r)
}

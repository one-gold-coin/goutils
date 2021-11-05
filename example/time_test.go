package main

import (
	"fmt"
	"github.com/one-gold-coin/goutils"
	"testing"
)

func Test_Now(t *testing.T) {
	fmt.Println(goutils.LocationAsiaSh+" now: ", goutils.Now())
	fmt.Println(goutils.LocationAmericaNewYork+" now: ", goutils.Now(goutils.LocationAmericaNewYork))
}

func Test_Int64Time(t *testing.T) {
	var unixTime int64
	unixTime = 1633774744

	fmt.Println(" Int64Time Default template: ", goutils.Int64Time(unixTime).String())
	fmt.Println(" Int64Time Default use template 1: ", goutils.Int64Time(unixTime).String(goutils.LayoutYmd))
	fmt.Println(" Int64Time Default use template 2: ", goutils.Int64Time(unixTime).String(goutils.LayoutHis))
	fmt.Println(" Int64Time Default int: ", goutils.Int64Time(unixTime).Int64())
}

func Test_StringTime(t *testing.T) {
	date := "2021-10-09 18:19:04"
	fmt.Println(" StringTime Default template: ", goutils.StringTime(date).Int64())
	fmt.Println(" StringTime Default use template 1: ", goutils.StringTime(date).Int64())
	fmt.Println(" StringTime Default use template 2: ", goutils.StringTime(date).Int64())
	fmt.Println(" StringTime Default string: ", goutils.StringTime(date).String())
}

func Test_DayStartUnix(t *testing.T) {
	fmt.Println(" DayStartUnix: ", goutils.DayStartUnix())
	fmt.Println(" DayEndUnix: ", goutils.DayEndUnix())
}

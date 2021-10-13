package goutils

import (
	"fmt"
	"testing"
)

func Test_Now(t *testing.T) {
	fmt.Println(LocationAsiaSh+" now: ", Now())
	fmt.Println(LocationAmericaNewYork+" now: ", Now(LocationAmericaNewYork))
}

func Test_Int64Time(t *testing.T) {
	var unixTime int64
	unixTime = 1633774744
	fmt.Println(" Int64Time Default template: ", Int64Time(unixTime).String())
	fmt.Println(" Int64Time Default use template 1: ", Int64Time(unixTime).String(TemplateYmd))
	fmt.Println(" Int64Time Default use template 2: ", Int64Time(unixTime).String(TemplateHis))
	fmt.Println(" Int64Time Default int: ", Int64Time(unixTime).Int64())
}

func Test_StringTime(t *testing.T) {
	date := "2021-10-09 18:19:04"
	fmt.Println(" StringTime Default template: ", StringTime(date).Int64())
	fmt.Println(" StringTime Default use template 1: ", StringTime(date).Int64(TemplateYmd))
	fmt.Println(" StringTime Default use template 2: ", StringTime(date).Int64(TemplateHis))
	fmt.Println(" StringTime Default string: ", StringTime(date).String())
}

func Test_DayStartUnix(t *testing.T) {
	fmt.Println(" DayStartUnix: ", DayStartUnix())
	fmt.Println(" DayEndUnix: ", DayEndUnix())
}

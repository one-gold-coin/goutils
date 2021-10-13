package goutils

import (
	"math/rand"
	"time"
)

func RandNum(minId, maxId int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(maxId-minId+1) + minId
}

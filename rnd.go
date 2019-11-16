package rnd

import (
	"math/rand"
	"time"
)

var (
	DataDir string
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

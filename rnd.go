package rnd

import (
	"math/rand"
	"time"
	"strings"
	"strconv"
	"fmt"
)

const alfabeto = "abcdefghijklmnopqrstuvwxyz"
const tamAlfabeto = 26

func init() {
	Seed()
}

func Seed() {
	str := strings.Replace(time.Now().Format("20060102150405.000"), ".", "", 1)
	ts, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		panic("rnd init: " + err.Error())
	}
	rand.Seed(ts)
}

func Hexa16() string {
	return fmt.Sprintf("%x", rand.Uint64())
}

func Hexa32() string {
	return fmt.Sprintf("%x%x", rand.Uint64(), rand.Uint64())
}

func Hexa64() string {
	return fmt.Sprintf("%x%x%x%x", rand.Uint64(), rand.Uint64(), rand.Uint64(), rand.Uint64())
}

func Alfa(tam int) string {
	var ret string
	for i := 0; i < tam; i++ {
		p := rand.Intn(tamAlfabeto)
		ret += alfabeto[p:p+1]
	}
	return ret
}

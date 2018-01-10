package rnd

import (
	"testing"
	"fmt"
)

func TestHexas(t *testing.T) {
	fmt.Println(Hexa16())
	fmt.Println(Hexa32())
	fmt.Println(Hexa64())
}

func TestAlfa(t *testing.T) {
	fmt.Println(Alfa(19))
}

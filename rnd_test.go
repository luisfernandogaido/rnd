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

//https://til.hashrocket.com/posts/355f31f19c-seeding-golangs-rand
//https://stackoverflow.com/questions/12321133/golang-random-number-generator-how-to-seed-properly
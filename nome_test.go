package rnd

import (
	"fmt"
	"testing"
)

func TestNome(t *testing.T) {
	for i := 0; i < 10; i++ {
		fmt.Println(Nome())
	}
}

func TestCidade(t *testing.T) {
	for i := 0; i < 10; i++ {
		fmt.Println(Cidade())
	}
}

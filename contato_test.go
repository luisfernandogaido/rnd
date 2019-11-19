package rnd

import (
	"fmt"
	"testing"
)

func TestEndereco(t *testing.T) {
	for i := 0; i < 10; i++ {
		fmt.Println(Endereco())
	}
}

func TestNumero(t *testing.T) {
	for i := 0; i < 10; i++ {
		fmt.Println(Numero())
	}
}

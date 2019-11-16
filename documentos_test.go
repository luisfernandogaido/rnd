package rnd

import (
	"fmt"
	"testing"
)

func TestCpf(t *testing.T) {
	for i := 0; i < 10; i++ {
		fmt.Println(Cpf())
	}
}

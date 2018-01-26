package rnd

import (
	"fmt"
	"testing"
)

func TestNewGeradorNomesFiles(t *testing.T) {
	gn, err := NewGeradorNomesFiles("./dados/nm.txt", "./dados/nf.txt", "./dados/sn.txt")
	if err != nil {
		t.Fatal(err)
	}
	for i := 0; i < 10; i++ {
		fmt.Println(gn.Nome(2, 3))
	}
}

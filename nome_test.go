package rnd

import (
	"fmt"
	"testing"
)

func TestLoadNomesTest(t *testing.T) {
	DataDir = "./dados"
	if err := LoadNomes(); err != nil {
		t.Fatal(err)
	}
}

func TestNome(t *testing.T) {
	DataDir = "./dados"
	if err := LoadNomes(); err != nil {
		t.Fatal(err)
	}
	for i := 0; i < 10; i++ {
		fmt.Println(Nome())
	}
}

package rnd

import (
	"fmt"
	"testing"
)

func TestEmail(t *testing.T) {
	if err := LoadNomes(); err != nil {
		t.Fatal(err)
	}
	for i := 0; i < 10; i++ {
		fmt.Println(Email(Nome()))
	}
}

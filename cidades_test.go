package rnd

import (
	"fmt"
	"testing"
)

func TestNewGeradorCidades(t *testing.T) {
	gc, err:= NewGeradorCidades("./dados/cidades.csv")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(gc.GeraNomeUf())
}
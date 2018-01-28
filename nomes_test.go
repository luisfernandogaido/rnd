package rnd

import (
	"fmt"
	"testing"
)

var gn *GeradorNomes

func init() {
	var err error
	gn, err = NewGeradorNomesFiles("./dados/nm.txt", "./dados/nf.txt", "./dados/sn.txt")
	if err != nil {
		panic(err)
	}
}

func TestEndereco(t *testing.T) {
	fmt.Println(Endereco(gn))
}

func TestNumeroEndereco(t *testing.T) {
	fmt.Println(NumeroEndereco())
}

func TestBairro(t *testing.T) {
	fmt.Println(Bairro(gn))
}

func TestEmail(t *testing.T) {
	fmt.Println(Email(gn.Nome(2, 2)))
}

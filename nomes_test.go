package rnd

import (
	"fmt"
	"io/ioutil"
	"strings"
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

func TestMaia(t *testing.T) {
	nomes := make([]string, 0)
	for i := 0; i < 1000000; i++ {
		nomes = append(nomes, gn.Nome(2, 3))
	}
	err := ioutil.WriteFile("./um-milhao-nomes.txt", []byte(strings.Join(nomes, "\r\n")), 0644)
	if err != nil {
		t.Fatal(err)
	}
}

package rnd

import (
	"fmt"
	"math/rand"
)

type nome struct {
	Nome   string
	Genero string
}

var (
	nomes         []nome
	tamNomes      int
	sobrenomes    []string
	tamSobrenomes int
)

func Nome() string {
	return fmt.Sprintf("%v %v", nomes[rand.Intn(tamNomes)].Nome, sobrenomes[rand.Intn(tamSobrenomes)])
}

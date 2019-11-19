package rnd

import (
	"fmt"
	"math/rand"
)

type nome struct {
	Nome   string
	Genero string
}

type cidade struct {
	nome string
	uf   string
}

var (
	nomes         []nome
	tamNomes      int
	sobrenomes    []string
	tamSobrenomes int
	cidades       []cidade
	tamCidade     int
)

func Nome() string {
	return fmt.Sprintf("%v %v", nomes[rand.Intn(tamNomes)].Nome, sobrenomes[rand.Intn(tamSobrenomes)])
}

func Cidade() string {
	c := cidades[rand.Intn(tamCidade)]
	return fmt.Sprintf("%v/%v", c.nome, c.uf)
}

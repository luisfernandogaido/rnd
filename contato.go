package rnd

import (
	"fmt"
	"math/rand"
	"strings"
)

var (
	dominios = [...]string{
		"@gmail.com",
		"@yahoo.com.br",
		"@hotmail.com",
		"@outlook.com",
		"@yopmail.com",
		"@uol.com.br",
		"@bol.com.br",
		"@globomail.com",
	}
	tamDominios = 8

	separadores = [...]string{
		".",
		"-",
		"_",
		"",
	}
	tamSeparadores = 4

	ruas = [...]string{
		"R.",
		"AV.",
		"TR.",
		"ROD.",
	}
	tamRuas = 4
)

func Email(nomeBase string) string {
	email := strings.ReplaceAll(strings.ToLower(nomeBase), " ", separadores[rand.Intn(tamSeparadores)])
	return email + dominios[rand.Intn(tamDominios)]
}

func Telefone() string {
	return Digitos(11)
}

func Cep() string {
	return Digitos(8)
}

func Endereco() string {
	return fmt.Sprintf("%v %v", ruas[rand.Intn(tamRuas)], Nome())
}

func Numero() string {
	return fmt.Sprintf("%v", rand.Intn(1000))
}

func Bairro() string {
	return Nome()
}

package rnd

import (
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
)

func Email(nomeBase string) string {
	email := strings.ReplaceAll(strings.ToLower(nomeBase), " ", separadores[rand.Intn(tamSeparadores)])
	return email + dominios[rand.Intn(tamDominios)]
}

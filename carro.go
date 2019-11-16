package rnd

import "strings"

func Placa() string {
	return strings.ToUpper(Letras(3)) + Digitos(4)
}

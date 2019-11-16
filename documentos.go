package rnd

import (
	"fmt"
	"math/rand"
	"strings"
)

var (
	digitos = "01234546789"
)

func Digitos(n int) string {
	sb := strings.Builder{}
	for i := 0; i < n; i++ {
		p := rand.Intn(10)
		sb.WriteString(digitos[p : p+1])
	}
	return sb.String()
}

func CpfH() string {
	return fmt.Sprintf("%v.%v.%v-%v", Digitos(3), Digitos(3), Digitos(3), Digitos(2))
}

func Cpf() string {
	return strings.ReplaceAll(strings.ReplaceAll(CpfH(), ".", ""), "-", "")
}

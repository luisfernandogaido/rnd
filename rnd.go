package rnd

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

const alfabeto = "abcdefghijklmnopqrstuvwxyz"
const tamAlfabeto = 26

var tiposLogradouro = []string{"Alameda", "Alameda", "Beco", "Estrada", "Praça", "Rua", "Via"}
var dominios = []string{"hotmail.com", "yahoo.com.br", "gmail.com", "uol.com.br", "bol.com.br"}
var separadores = []string{"", ".", "_", "-"}

func init() {
	Seed()
}

func Seed() {
	str := strings.Replace(time.Now().Format("20060102150405.000"), ".", "", 1)
	ts, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		panic("rnd init: " + err.Error())
	}
	rand.Seed(ts)
}

func Hexa16() string {
	return fmt.Sprintf("%x", rand.Uint64())
}

func Hexa32() string {
	return fmt.Sprintf("%x%x", rand.Uint64(), rand.Uint64())
}

func Hexa64() string {
	return fmt.Sprintf("%x%x%x%x", rand.Uint64(), rand.Uint64(), rand.Uint64(), rand.Uint64())
}

func Alfa(tam int) string {
	var ret string
	for i := 0; i < tam; i++ {
		p := rand.Intn(tamAlfabeto)
		ret += alfabeto[p : p+1]
	}
	return ret
}

func Endereco(gn *GeradorNomes) string {
	return sorteiaString(tiposLogradouro) + " " + gn.Nome(2, 1)
}

func NumeroEndereco() string {
	if rand.Float32() > 0.75 {
		return strconv.Itoa(rand.Intn(99)+1) + "-" + strconv.Itoa(rand.Intn(199)+1)
	} else {
		return strconv.Itoa(rand.Intn(9998) + 1)
	}
}

func Bairro(gn *GeradorNomes) string {
	return gn.Nome(1, 1, )
}

func Email(nome string) string {
	email := strings.ToLower(nome) + "@" + sorteiaString(dominios)
	separador := sorteiaString(separadores)
	replacer := strings.NewReplacer(
		" ", separador,
		"/", separador,
		",", "",
		"á", "a",
		"é", "e",
		"í", "i",
		"ó", "o",
		"ú", "u",
		"â", "a",
		"ê", "e",
		"î", "i",
		"ô", "o",
		"û", "u",
		"ã", "a",
		"õ", "o",
		"ç", "c",
		"à", "a",
		"è", "e",
		"ì", "i",
		"ò", "o",
		"ù", "u",
		"'", "",
		"?", "",
		"ø", "o",
		"(", "",
		")", "",
	)
	return replacer.Replace(email)
}

func Cpf() string {
	return digitos(3) + "." + digitos(3) + "." + digitos(3) + "-" + digitos(2)
}

func Cep() string {
	return digitos(5) + "-" + digitos(3)
}

func sorteiaString(slice []string) string {
	return slice[rand.Intn(len(slice))]
}

func digitos(n int) string {
	var ret string
	for i := 0; i < n; i++ {
		ret += strconv.Itoa(rand.Intn(10))
	}
	return ret
}

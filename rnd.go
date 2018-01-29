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

// Gera um CPF. Com f true, formata.
func Cpf(f bool) string {
	if f {
		return digitos(3) + "." + digitos(3) + "." + digitos(3) + "-" + digitos(2)
	}
	return digitos(11)
}

// Gera um CEP. Com f true, formata.
func Cep(f bool) string {
	if f {
		return digitos(5) + "-" + digitos(3)
	}
	return digitos(8)
}

//Gera uma data de nascimento para alguém que tem atualmente entre 18 e 70 anos, aproximadamente.
// Com f true, formata.
func Nascimento(f bool) string {
	hoje := time.Now()
	nascimento := hoje.Add(-time.Hour * time.Duration(157680+rand.Intn(455520)))
	if f {
		return nascimento.Format("02/01/2006")
	}
	return nascimento.Format("2006-01-02")
}

//Gera um RG. Com f true, formata
func Rg(f bool) string {
	if f {
		return digitos(2) + "." + digitos(3) + "." + digitos(3) + "-" + digitos(1)
	}
	return digitos(9)
}

//Gera um telefone. Com f true, formata.
func Telefone(f bool) string {
	if f {
		return fmt.Sprintf("(%v)%v-%v", digitos(2), digitos(4), digitos(4))
	}
	return digitos(11)
}

//Gera um CNPJ. Com f true, formata.
func Cnpj(f bool) string {
	if f {
		return fmt.Sprintf("%v.%v.%v/%v-%v", digitos(2), digitos(3), digitos(3), digitos(4), digitos(2))
	}
	return digitos(14)
}

// Dado um número, retorna um novo diferente em até, no máximo a taxa de desvio. n entre 0.0 e 1.0.
// Se desv não estiver entre 0.0 e 1.0, 1.0 é assumido
// Exemplos:
// n = 1.0 e desv = 0.2 resulta entre 0.8 e 1.2
// n = 1.0 e desv = 1.0 resulta entre 0.0 e 2.0
func RealModificado(n float64, desv float64) float64 {
	if desv < 0.0 || desv > 1 {
		desv = 1.0
	}
	sinal := 1.0
	if rand.Intn(2) == 1 {
		sinal = -1.0
	}
	return n + sinal*n*desv*rand.Float64()
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

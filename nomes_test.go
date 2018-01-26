package rnd

import (
	"testing"
	"io/ioutil"
	"strings"
	"fmt"
)

func Test(t *testing.T) {
	mapaPalavras := make(map[string]string)
	arquivos := []string{
		"./dados/nomes-masculinos.txt",
		"./dados/nomes-masculinos-old.txt",
		"./dados/nomes-femininos.txt",
		"./dados/nomes-femininos-old.txt",
		"./dados/sobrenomes.txt",
		"./dados/sobrenomes-tiltedlogic.txt",
		"./dados/CSV_Database_of_Last_Names.csv",
	}
	for _, arq := range arquivos {
		nomes, err := carregaArquivo(arq)
		if err != nil {
			t.Error(err)
		}
		var indice string
		for _, nome := range nomes {
			fields := strings.Fields(nome)
			for _, field := range fields {
				indice = strings.TrimSpace(strings.ToLower(field))
				mapaPalavras[indice] = field
			}
		}
		fmt.Println(len(mapaPalavras))
	}
	mapaPalavrasCorreios := make(map[string]string)
	nomes, err := carregaArquivo("./dados/funcionarios-correios.txt")
	if err != nil {
		t.Error(err)
	}
	var indice string
	for _, nome := range nomes {
		fields := strings.Fields(nome)
		for _, field := range fields {
			indice = strings.TrimSpace(strings.ToLower(field))
			mapaPalavrasCorreios[indice] = field
		}
	}
	for k, v := range mapaPalavrasCorreios {
		if _, ok := mapaPalavras[k]; !ok {
			fmt.Println(v)
		}
	}
}

func carregaArquivo(arq string) ([]string, error) {
	bytes, err := ioutil.ReadFile(arq)
	if err != nil {
		return nil, err
	}
	nomes := strings.Split(string(bytes), "\r\n")
	return nomes, nil
}

func une(mapa map[string]string, nomes []string) {
	for _, nome := range nomes {
		i := strings.TrimSpace(strings.ToLower(nome))
		mapa[i] = nome
	}
}

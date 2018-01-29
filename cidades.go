package rnd

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"strings"

	"github.com/pkg/errors"
)

type cidade struct {
	Nome, Estado, Uf string
}

type GeradorCidades struct {
	Cidades []cidade
}

// Cria um gerador de cidades com base em arquivo CSV.
// O arquivo deve conter em cada linha: <nome>\t<estado>\t<UF>\n.
// Se uma linha não satisfizer esse conteúdo, um erro é gerado.
func NewGeradorCidades(arq string) (*GeradorCidades, error) {
	cidades := make([]cidade, 0)
	bytes, err := ioutil.ReadFile(arq)
	if err != nil {
		return nil, err
	}
	linhas := strings.Split(string(bytes), "\n")
	for _, l := range linhas {
		linha := strings.Split(l, "\t")
		if len(linha) != 3 {
			return nil, errors.New("todas as linhas devem conter 3 campos: cidade, estado e uf")
		}
		cidades = append(cidades, cidade{linha[0], linha[1], linha[2]})
	}
	return &GeradorCidades{cidades}, nil
}

func (gc *GeradorCidades) GeraNomeUf() string {
	cidade := gc.Cidades[rand.Intn(len(gc.Cidades))]
	return fmt.Sprintf("%v/%v", cidade.Nome, cidade.Uf)
}

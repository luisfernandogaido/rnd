package rnd

import (
	"errors"
	"io/ioutil"
	"math/rand"
	"strings"
)

//Armazena slices de nomes masculinos, femininos e sobrenomes para gerar combinações de nomes.
type GeradorNomes struct {
	masculinos, femininos, sobrenomes []string
}

// Retorna um nome masculino com a quantidade máxima de nomes e sobrenomes.
// Se maxN e/ou maxSn for <= 0, assumem, respectivamente, os valores 2 e/ou 1.
func (g *GeradorNomes) Masculino(maxN, maxSn int) string {
	n := make([]string, 0)
	if maxN <= 0 {
		maxN = 2
	}
	if maxSn <= 0 {
		maxSn = 1
	}
	mn := rand.Intn(maxN) + 1
	msn := rand.Intn(maxSn) + 1
	for i := 0; i < mn; i++ {
		n = append(n, g.masculinos[rand.Intn(len(g.masculinos))])
	}
	for i := 0; i < msn; i++ {
		n = append(n, g.sobrenomes[rand.Intn(len(g.sobrenomes))])
	}
	return strings.Join(n, " ")
}

// Retorna um nome feminino com a quantidade máxima de nomes e sobrenomes.
// Se maxN e/ou maxSn for <= 0, assumem, respectivamente, os valores 2 e/ou 1.
func (g *GeradorNomes) Feminino(maxN, maxSn int) string {
	n := make([]string, 0)
	if maxN <= 0 {
		maxN = 2
	}
	if maxSn <= 0 {
		maxSn = 1
	}
	mn := rand.Intn(maxN) + 1
	msn := rand.Intn(maxSn) + 1
	for i := 0; i < mn; i++ {
		n = append(n, g.femininos[rand.Intn(len(g.femininos))])
	}
	for i := 0; i < msn; i++ {
		n = append(n, g.sobrenomes[rand.Intn(len(g.sobrenomes))])
	}
	return strings.Join(n, " ")
}

// Retorna um nome masculino ou feminino com a quantidade máxima de nomes e sobrenomes.
// Se maxN e/ou maxSn for <= 0, assumem, respectivamente, os valores 2 e/ou 1.
func (g *GeradorNomes) Nome(maxN, maxSn int) string {
	genero := rand.Intn(2)
	if genero == 0 {
		return g.Masculino(maxN, maxSn)
	} else {
		return g.Feminino(maxN, maxSn)
	}
}

// Constrói um gerador de nomes com slices de nomes masculinos, femininos e sobrenomes.
// Remove espaços adicionais.
func NewGeradorNomes(m, f, s []string) (*GeradorNomes, error) {
	if m == nil || f == nil || s == nil {
		return nil, errors.New("gerador de nomes: as três slices precisam estar inicializadas")
	}
	return &GeradorNomes{limpaEspacosNomes(m), limpaEspacosNomes(f), limpaEspacosNomes(s)}, nil
}

// Constrói um gerador de nomes com arquivos de nomes masculinos, femininos e sobrenomes.
// Os arquivos devem conter um nome por linha. As linhas devem ser quebradas com "\n".
// Remove espaços adicionais.
func NewGeradorNomesFiles(m, f, s string) (*GeradorNomes, error) {
	bytes, err := ioutil.ReadFile(m)
	if err != nil {
		return nil, errors.New("gerador de nomes: erro ao ler nomes masculinos")
	}
	nm := strings.Split(string(bytes), "\n")
	bytes, err = ioutil.ReadFile(f)
	if err != nil {
		return nil, errors.New("gerador de nomes: erro ao ler nomes femininos")
	}
	nf := strings.Split(string(bytes), "\n")
	bytes, err = ioutil.ReadFile(s)
	if err != nil {
		return nil, errors.New("gerador de nomes: erro ao ler sobrenomes")
	}
	sn := strings.Split(string(bytes), "\n")
	gn, err := NewGeradorNomes(nm, nf, sn)
	if err != nil {
		return nil, err
	}
	return gn, nil
}

func limpaEspacosNomes(nomes []string) []string {
	ret := make([]string, 0)
	for _, n := range nomes {
		ret = append(ret, strings.TrimSpace(n))
	}
	return ret
}

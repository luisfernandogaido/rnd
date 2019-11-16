package rnd

import (
	"encoding/csv"
	"fmt"
	"math/rand"
	"strings"
)

type nome struct {
	Nome   string
	Genero string
}

var (
	nomes         []nome
	tamNomes      int
	sobrenomes    []string
	tamSobrenomes int
)

func LoadNomes() error {
	if nomes != nil {
		return nil
	}
	r := csv.NewReader(strings.NewReader(nomesCsv))
	records, err := r.ReadAll()
	if err != nil {
		return fmt.Errorf("load nomes: %w", err)
	}
	records = records[1:]
	tamNomes = len(records)
	nomes = make([]nome, tamNomes)
	for i := range records {
		nomes[i].Nome = records[i][2]
		nomes[i].Genero = records[i][1]
	}
	sobrenomes = strings.Split(sobrenomesCsv, "\n")
	tamSobrenomes = len(sobrenomes)
	return nil
}

func Nome() string {
	return fmt.Sprintf("%v %v", nomes[rand.Intn(tamNomes)].Nome, sobrenomes[rand.Intn(tamSobrenomes)])
}

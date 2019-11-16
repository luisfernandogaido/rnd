package rnd

import (
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"path/filepath"
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
	f, err := os.Open(filepath.Join(DataDir, "nomes.csv"))
	if err != nil {
		return fmt.Errorf("load nomes: %w", err)
	}
	defer f.Close()
	r := csv.NewReader(f)
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
	b, err := ioutil.ReadFile(filepath.Join(DataDir, "sobrenomes.txt"))
	sobrenomes = strings.Split(string(b), "\r\n")
	tamSobrenomes = len(sobrenomes)
	return nil
}

func Nome() string {
	return fmt.Sprintf("%v %v", nomes[rand.Intn(tamNomes)].Nome, sobrenomes[rand.Intn(tamSobrenomes)])
}

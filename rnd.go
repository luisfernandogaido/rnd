package rnd

import (
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
	os.Setenv("RND_DIR", "C:\\GoPrograms\\rnd")
	if os.Getenv("RND_DIR") == "" {
		panic("rnd: RND_DIR não definida")
	}
	f, err := os.Open(filepath.Join(os.Getenv("RND_DIR"), "nomes.csv"))
	if err != nil {
		panic(fmt.Errorf("load nomes: %w", err))
	}
	defer f.Close()
	r := csv.NewReader(f)
	records, err := r.ReadAll()
	if err != nil {
		panic(fmt.Errorf("load nomes: %w", err))
	}
	records = records[1:]
	tamNomes = len(records)
	nomes = make([]nome, tamNomes)
	for i := range records {
		nomes[i].Nome = records[i][2]
		nomes[i].Genero = records[i][1]
	}
	b, err := ioutil.ReadFile(filepath.Join(os.Getenv("RND_DIR"), "sobrenomes.txt"))
	sobrenomes = strings.Split(string(b), "\r\n")
	tamSobrenomes = len(sobrenomes)
}

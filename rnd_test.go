package rnd

import (
	"fmt"
	"testing"
)

func TestHexas(t *testing.T) {
	fmt.Println(Hexa16())
	fmt.Println(Hexa32())
	fmt.Println(Hexa64())
}

func TestAlfa(t *testing.T) {
	fmt.Println(Alfa(19))
}

func TestNascimento(t *testing.T) {
	fmt.Println(Nascimento(true))
	fmt.Println(Nascimento(false))
}

func TestRg(t *testing.T) {
	fmt.Println(Rg(true))
	fmt.Println(Rg(false))
}

func TestCpf(t *testing.T) {
	fmt.Println(Cpf(true))
	fmt.Println(Cpf(false))
}

func TestCep(t *testing.T) {
	fmt.Println(Cep(true))
	fmt.Println(Cep(false))
}

func TestTelefone(t *testing.T) {
	fmt.Println(Telefone(true))
	fmt.Println(Telefone(false))
}

func TestCnpj(t *testing.T) {
	fmt.Println(Cnpj(true))
	fmt.Println(Cnpj(false))
}

func TestModificaReal(t *testing.T) {
	fmt.Println(RealModificado(1.0, 0.2))
	fmt.Println(RealModificado(1.0, 0.2))
	fmt.Println(RealModificado(1.0, 0.2))
	fmt.Println(RealModificado(1.0, 0.2))
}

//https://til.hashrocket.com/posts/355f31f19c-seeding-golangs-rand
//https://stackoverflow.com/questions/12321133/golang-random-number-generator-how-to-seed-properly

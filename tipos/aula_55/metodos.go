package main

import (
	"fmt"
	"strings"
)

type pessoa struct {
	nome      string
	sobrenome string
}

func (p pessoa) getNomeCompleto() string{
	// p está associado a estrutura de dados pessoa. Por isso implementa o método.
	return p.nome + " " + p.sobrenome
}

func (p *pessoa) setNomeCompleto(nomeCompleto string) {
	partes := strings.Split(nomeCompleto, " ")
	p.nome = partes[0]
	p.sobrenome = partes[1]
}

func main() {
	p1 := pessoa{nome: "João", sobrenome: "Silva"}
	fmt.Println(p1.getNomeCompleto())

	p1.setNomeCompleto("Maria Conceição")
	fmt.Println(p1.getNomeCompleto())

}

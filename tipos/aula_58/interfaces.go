package main

import "fmt"

/*
Go trouxe o método e interface como inspiração de OO.

Go aplica interfaces de forma diferente.
Em demais linguagens OO, é necessário declarar explicitamente dentro
do objeto o que implementa determinada interface.

Em Go, se uma estrutura respeita os métodos de uma interface, o pŕoprio compilador
infere identificando que a estrutura tem método que respeitam os métodos da interface,
sendo assim ela é do tipo.

O polimorfismo funciona sem ter que explicitamente amarrar na struct.
*/

type imprimivel interface {
	// Se existe uma estrutura que respeita
	toString() string
}

type pessoa struct {
	nome      string
	sobrenome string
}

type produto struct {
	nome  string
	preco float64
}

// interface são implementadas implicitamente
func (p pessoa) toString() string {
	return p.nome + " " + p.sobrenome
}

func (p produto) toString() string {
	return fmt.Sprintf("%s - R$ %.2f", p.nome, p.preco)
}

func imprimir(x imprimivel) {
	fmt.Println(x.toString())
}

func main() {
	var coisa imprimivel = pessoa{"Roberto", "Soares"}
	fmt.Println(coisa.toString())
	imprimir(coisa)

	prd := produto{"Leite Condesado", 5.19}
	fmt.Println(prd.toString())

}

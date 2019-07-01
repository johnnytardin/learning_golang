package main

import (
	"fmt"
)

type produto struct {
	nome     string
	preco    float64
	desconto float64
}

// Método: função com receiver (receptor)
func (p produto) precoComDesconto() float64 {
	return p.preco * (1 - p.desconto)
}

func main() {
	var produto1 produto
	produto1 = produto{
		nome:     "Ferrari California",
		preco:    1370000.00,
		desconto: 0.02,
	}
	fmt.Println(produto1, produto1.precoComDesconto())

	produto2 := produto{"Lapis", 2.10, 0.06}
	fmt.Println(produto2, produto2.precoComDesconto())

}

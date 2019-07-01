package main

import "fmt"

type item struct {
	produtoID int
	qtde      int
	preco     float64
}

type pedido struct {
	userID int
	itens  []item
}

func (p pedido) valorTotal() float64 {
	total := 0.0
	for _, item := range p.itens {
		total += item.preco * float64(item.qtde)
	}
	return total
}

func main() {
	pedido := pedido{
		userID: 1000012,
		itens: []item{
			item{produtoID: 101, qtde: 2, preco: 100.21},
			item{produtoID: 109, qtde: 1, preco: 90},
			item{produtoID: 123, qtde: 12, preco: 8.99},
		},
	}
	fmt.Printf("O valor total do pedido é de R$ %.2f", pedido.valorTotal())
}

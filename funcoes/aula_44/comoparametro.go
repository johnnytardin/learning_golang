package main

import "fmt"

func multiplicacao(a, b int) int {
	return a * b
}

func exec(funcao func(int, int) int, p1, p2 int) int {
	/*
		map, endorse, filter são funções para estudar implemetação (Go não tem)
		de função como parâmetro.
	*/
	return funcao(p1, p2)
}

func main() {
	resultado := exec(multiplicacao, 9, 6)
	fmt.Println(resultado)
}

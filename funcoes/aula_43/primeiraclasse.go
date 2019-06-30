package main

import "fmt"

var soma = func(a, b int) int {
	// função anônima
	return a + b
}

func main() {
	fmt.Println(soma(9, 1))

	// função armazenada em uma variável (paradigma funcional)
	sub := func(a, b int) int {
		return a - b
	}
	fmt.Println(sub(2, 20))
}

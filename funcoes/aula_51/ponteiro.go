package main

import "fmt"

func inc1(n int) {
	n++
}

// revisão: um ponteiro é representado por *
func inc2(n *int) {
	// revisão: * é usado para desreferenciar, ou seja,
	// ter acesso ao qual o ponteiro aponta.
	*n++
	// ponteiros não fazem operações atitméticas, por isso foi preciso,
	// desreferenciar para então incrementar.
}

func main() {
	n := 1

	inc1(n) // por valor
	fmt.Println(n)

	// revisão: & é usado para obter o endereço da variável
	inc2(&n) // por referência
	fmt.Println(n)

	// evitar variáveis com referências, pois torna o código confuso e
	// passível de bugs, priorizando passar o valor para as funções.

}

package main

import "fmt"

func fatorial(n uint) uint {
	/*
		Simplificando a aula anterior, basta passar o tipo de patâmetro de int para uint
		pois uint não contém números negativos.
		Caso seja feita a tentativa de usar um número negativo setando diretamento no código,
		não será possível compilar.
	*/

	// Multiplica pelo número anterior até chegar a 1
	switch {
	case n == 0:
		// condição de parada da fatorial
		return 1
	default:
		fatorialAnterior := fatorial(n - 1)
		return n * fatorialAnterior

	}
}

func main() {
	for _, numero := range []uint{5} {
		resultado := fatorial(numero)
		fmt.Println(resultado)

	}
}

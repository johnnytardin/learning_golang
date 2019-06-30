package main

import "fmt"

func fatorial(n int) (int, error) {
	// Multiplica pelo número anterior até chegar a 1
	switch {
	case n < 0:
		return -1, fmt.Errorf("número inválido %d", n)
	case n == 0:
		return 1, nil
	default:
		fatorialAnterior, _ := fatorial(n - 1)
		return n * fatorialAnterior, nil

	}
}

func main() {
	resultado, _ := fatorial(10)
	fmt.Println(resultado)
}

package main

import "fmt"

func main() {
	// var aprovados map[int]string
	// mapas devem ser inicializados
	aprovados := make(map[int]string)

	aprovados[1234567890] = "João"
	aprovados[1987654321] = "Pedro"
	aprovados[1234567891] = "Ana"

	for cpf, nome := range aprovados {
		fmt.Printf("%s (CPF: %d)\n", nome, cpf)
	}

	fmt.Println(aprovados[1234567890])
	delete(aprovados, 1234567890)
	fmt.Println(aprovados[1234567890])

}

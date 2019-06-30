package main

import "fmt"

func imprimirAprovados(aprovados ...string) {
	fmt.Println("Lista de Aprovados")

	for i, aprovado := range aprovados {
		fmt.Printf("%d) %s\n", i+1, aprovado)
	}
}

func main() {
	aprovados := []string{"Maria", "Pedro", "Rafael", "Guilherme"}

	// Passando um slice para função variática: basta passar com os ... no final.
	// Não funcionaria com array [3]string{}
	imprimirAprovados(aprovados...)

}

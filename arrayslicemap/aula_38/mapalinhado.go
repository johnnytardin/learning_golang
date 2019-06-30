package main

import "fmt"

func main() {
	funcsPorLetra := map[string]map[string]float64{
		"G": {
			"Gabriela Silva": 15000.0,
			"Guga Ferreira":  18000.00,
		},
		"J": {
			"José": 1200.00,
			"João": 1456.00,
		},
		"P": {
			"Pedro": 100000.00,
		},
	}

	delete(funcsPorLetra, "J")

	for letra, funcs := range funcsPorLetra {
		fmt.Printf("Funcionários com a letra: %s\n", letra)
		for nome, salario := range funcs {
			fmt.Printf("%s com salário de %.2f\n", nome, salario)
		}
	}
}

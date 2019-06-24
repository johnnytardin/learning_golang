package main

import "fmt"

func imprimitResultado(nota float64) {
	if nota >= 7 {
		fmt.Println("Aprovado com nota", nota)
	} else {
		fmt.Println("Reprovado com nota", nota)
	}
}

func main() {
	imprimitResultado(7.3)
	imprimitResultado(5.1)
}
package main

import "fmt"

func obterValorAprovado(numero int) int {
	// defer atrasa a execução dessa linha 7
	defer fmt.Println("fim!")
	if numero > 5000 {
		fmt.Println("muito alto")
		return 5000
	} 
	fmt.Println("Valor baixo...")
	return numero

}

func main() {
	fmt.Println(obterValorAprovado(6000))
}
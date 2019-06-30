package main

import "fmt"

func closure() func() {
	x := 10
	var funcao = func() {
		fmt.Println(x)
	}
	return funcao
}

func main() {
	x := 20

	closureX := closure()
	closureX()
	fmt.Println("No contexto main x é:", x)
}

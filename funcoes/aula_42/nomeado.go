package main

import "fmt"

func trocar(p1, p2 int) (segundo, primeiro int) {
	primeiro = p1
	segundo = p2
	return // retorno limpo. como acima os valores de retorno já foram atribuidos.
}

func main() {
	r1, r2 := trocar(9, 0)
	fmt.Println(r1, r2)
}

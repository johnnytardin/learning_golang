package main

import "fmt"

func main() {
	p1 := Ponto{2.0, 11.0}
	p2 := Ponto{3.0, 17.0}

	fmt.Println(catetos(p1, p2))
	fmt.Println(Distancia(p1, p2))

}

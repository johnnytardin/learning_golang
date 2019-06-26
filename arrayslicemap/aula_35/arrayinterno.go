package main

import "fmt"

func main() {
	s1 := make([]int, 10, 20)
	s2 := append(s1, 1, 2, 3)
	fmt.Println(s1, s2)

	s1[0] = 7
	/* Mudando a posicao 1 do s1, automaticamente s2 tambem tem o mesmo valor.
	Com isso fica provado que s2 aponta para o mesmo objeto.
	*/
	fmt.Println(s1, s2)

}

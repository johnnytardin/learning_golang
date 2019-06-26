package main

import "fmt"

func main() {
	s := make([]int, 10)
	s[9] = 12
	fmt.Println(s)

	s = make([]int, 10, 20) // crie um novo slice de tamanho 10, com capacidade de 20.
	fmt.Println(s, len(s), cap(s))

	s = append(s, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0) // incremente valores ao slice
	fmt.Println(s, len(s), cap(s))

	s = append(s, 1) // slice atingiu seu tamanho maximo (20), com isso o Go incrementou a capacidade do slice com mais 20.
	fmt.Println(s, len(s), cap(s))

}

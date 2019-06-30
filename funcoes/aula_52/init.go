package main

import "fmt"

func init() {
	// Por convenção, a função init é iniciada antes de qualquer outra.
	fmt.Println("Iniciando...")
}

func main() {
	fmt.Println("Main...")
}

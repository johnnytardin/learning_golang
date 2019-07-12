package main

import (
	"fmt"
	"time"
)

func fale(pessoa, texto string, qtde int) {
	time.Sleep(time.Second)
	for i := 0; i < qtde; i++ {
		fmt.Printf("%s: %s (iteração %d)\n", pessoa, texto, i+1)
	}
}

func main() {
	// fale("Maria", "Ei...", 5)
	// fale("João", "Opa...", 3)

	// como está na main se somente o código abaixo for executado, a função main 
	// vai finalizar antes da execução da função fale, e as goroutines não chegam
	// a ser executadas até o fim.
	go fale("Gerson", "Epaaaa", 200)
	fale("José", "hahahahaah", 100)


}

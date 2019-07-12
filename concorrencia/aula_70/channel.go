package main

import (
	"fmt"
)

func main() {
	ch := make(chan int, 1)

	ch <- 1 // enviando dados para o canal
	<-ch    // recebendo dados de um canal

	ch <- 2

	fmt.Println(<-ch)
}

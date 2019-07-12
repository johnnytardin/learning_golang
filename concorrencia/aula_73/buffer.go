package main

import (
	"fmt"
	"time"
)

/*
O bloqueio acontece quando o buffer chega em seu limite
*/

func rotina(ch chan int) {
	ch <- 1
	fmt.Println("Passei pelo 1")
	ch <- 2
	fmt.Println("Passei pelo 2")
	ch <- 3
	fmt.Println("Passei pelo 3")
	ch <- 4 // nesse caso vai gerar bloqueio aqui pois o buffer tem tamanho 3
	fmt.Println("Passei pelo 4")
	ch <- 5
	fmt.Println("Passei pelo 5")
	ch <- 6
	fmt.Println("Passei pelo 6")

}

func main() {
	ch := make(chan int, 3) // o segundo parâmetro com tamanho 3
	go rotina(ch)

	time.Sleep(time.Second)
	fmt.Println(<-ch)
}

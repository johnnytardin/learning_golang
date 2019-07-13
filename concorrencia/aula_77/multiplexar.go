package main

import (
	"fmt"

	"github.com/johnnytardin/html"
)

func encaminhar(origem <-chan string, destino chan string) {
	for {
		destino <- <-origem
	}
}

func multiplexar(entrada1, entrada2 <-chan string) <-chan string {
	c := make(chan string)
	go encaminhar(entrada1, c)
	go encaminhar(entrada2, c)
	return c
}

func main() {
	c := multiplexar(
		html.Titulo("https://www.google.com", "https://www.uol.com.br"),
		html.Titulo("https://www.icq.com", "https://www.msn.com"),
	)
	fmt.Println(<-c, "|", <-c)
	fmt.Println(<-c, "|", <-c)
}

package main

import (
	"fmt"
)

func falar(pessoa string) <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; i < 3; i++ {
			c <- fmt.Sprintf("%s falando: %d", pessoa, i)
		}

	}()
	return c
}

func multiplexar(entrada1, entrada2 <-chan string) <-chan string {
	c := make(chan string)
	go func() {
		for {
			select {
			case s := <-entrada1:
				c <- s
			case s := <-entrada2:
				c <- s
			}
		}
	}()
	return c
}

func main() {
	c := multiplexar(falar("João"), falar("Maria"))
	// aqui espera sempre causando o bloqueio pra seguir na próxima linha
	fmt.Println(<-c, <-c)
	fmt.Println(<-c, <-c)
	fmt.Println(<-c, <-c)
}

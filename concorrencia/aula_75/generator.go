package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

// https://www.youtube.com/watch?v=f6kdp27TYZs

// <-chan - canal somente-leitura
func titulo(urls ...string) <-chan string {
	// Aqui a função retorna um canal, com os resultados.
	// Com isso o cliente não sabe da estrutura interna encapsulando as goroutines que são executudas de forma concorrente.
	c := make(chan string)
	for _, url := range urls {
		go func(url string) {
			resp, _ := http.Get(url)
			html, _ := ioutil.ReadAll(resp.Body)

			r, _ := regexp.Compile("<title>(.*?)<\\/title>")
			c <- r.FindStringSubmatch(string(html))[1]
		}(url)
	}
	// o channel é retornado já pronto
	return c
}

func main() {
	// Aqui está uma mudança original do curso onde uso um slice passando para a função variática.
	t1 := titulo([]string{"https://www.yahoo.com.br", "https://www.uol.com.br"}...)
	t2 := titulo([]string{"https://www.microsoft.com", "http://www.google.com"}...)

	fmt.Println("Primeiros:", <-t1, "|", <-t2)
	fmt.Println("Segundos:", <-t1, "|", <-t2)
}
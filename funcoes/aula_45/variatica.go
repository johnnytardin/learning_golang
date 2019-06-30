package main

import (
	"fmt"
)

func media(numeros ...float64) float64 {
	total := 0.0
	for _, numero := range numeros {
		total += numero
	}
	return total / float64(len(numeros))

}

func main() {
	fmt.Printf("Média %.2f", media(5, 77.8, 9, 5.8, 4.6, 32.7))
}

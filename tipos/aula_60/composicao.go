package main

import "fmt"

type esportivo interface {
	ligarTurbo()
}

type ferrari struct {
	modelo          string
	turboLigado     bool
	velocidadeAtual int
}

func (f *ferrari) ligarTurbo() {
	f.turboLigado = true
}

func main() {
	c1 := ferrari{"F40", false, 0}
	c1.ligarTurbo()

	var c2 esportivo = &ferrari{"f40", false, 0}
	c2.ligarTurbo()

	fmt.Println(c1, c2)

}

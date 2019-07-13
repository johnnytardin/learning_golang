package matematica

import (
	"testing"
)

const erroPadrao = "Valor esperado %v, mas o resultado encontrado foi %v."

// teste unitário é testar uma unidade, quanto mais isolado do mundo exterior, o teste é melhor pois é visto de forma única.

func TestMedia(t *testing.T) {
	valorEsperado := 7.28
	valor := Media(7.2, 9.9, 6.1, 5.9)

	if valor != valorEsperado {
		t.Errorf(erroPadrao, valorEsperado, valor)
	}
}

// É possível acessar a pasta com o comando:
// go test

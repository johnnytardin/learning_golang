package arquitetura

import (
	"runtime"
	"testing"
)

func TestDependente(t *testing.T) {
	if runtime.GOARCH == "amd64" {
		t.Skip("Não funciona em arquitetura amd64")
	}
	t.Fail()

	// go test std - vai executar os testes nas standard libraries
	// t.Parallel() - roda os testes de forma papalela com outros testes

}

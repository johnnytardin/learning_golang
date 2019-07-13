package matematica

import (
	"fmt"
	"strconv"
)

// Media é responsável por fazer o que é obvio =)
func Media(numeros ...float64) float64 {
	total := 0.0
	for _, n := range numeros {
		total += n
	}
	media := total / float64(len(numeros))
	mediaArredondada, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", media), 64)
	return mediaArredondada
}

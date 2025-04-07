package main

import (
	"fmt"
	"math"
)

type Competidor struct {
	pedraA    int
	pedraB    int
	eliminado bool
}

func NewCompetidor(pedraA, pedraB int) Competidor {
	competidor := Competidor{pedraA: pedraA, pedraB: pedraB}
	if pedraA < 10 || pedraB < 10 {
		competidor.eliminado = true
	}
	return competidor
}

func (c Competidor) Pontuacao() int {
	return int(math.Abs(float64(c.pedraA - c.pedraB)))
}

func main() {
	var n int
	fmt.Scan(&n)

	melhorIndice := -1
	melhorPontuacao := 101

	for i := range n {
		var a, b int
		fmt.Scan(&a, &b)

		competidor := NewCompetidor(a, b)

		if !competidor.eliminado {
			pontuacao := competidor.Pontuacao()
			if pontuacao < melhorPontuacao {
				melhorPontuacao = pontuacao
				melhorIndice = i
			}
		}
	}

	if melhorIndice == -1 {
		fmt.Println("sem ganhador")
	} else {
		fmt.Println(melhorIndice)
	}
}

package main

import "fmt"

type Pos struct {
	x, y int
}

func atualizar_cobra(cobra []Pos, direcao string) {
	anterior := cobra[0]

	switch direcao {
	case "L":
		cobra[0].x--
	case "R":
		cobra[0].x++
	case "U":
		cobra[0].y--
	case "D":
		cobra[0].y++
	}

	for i := 1; i < len(cobra); i++ {
		temp := cobra[i]
		cobra[i] = anterior
		anterior = temp
	}
}

func main() {
	var q int
	var d string
	fmt.Scan(&q, &d)

	cobra := make([]Pos, q)

	for i := range q {
		fmt.Scan(&cobra[i].x, &cobra[i].y)
	}

	atualizar_cobra(cobra, d)

	for i := range q {
		fmt.Println(cobra[i].x, cobra[i].y)
	}
}

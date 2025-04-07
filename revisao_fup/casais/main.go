package main

import "fmt"

func main() {

	var n int
	fmt.Scan(&n)

	vetor := make([]int, n)

	for i := range vetor {
		fmt.Scan(&vetor[i])
	}

	pares := 0
	for i := range vetor {
		for j := i + 1; j < n; j++ {
			if vetor[i] == -vetor[j] && vetor[i] != 0 && vetor[j] != 0 {
				vetor[i] = 0
				vetor[j] = 0
				pares++
				break
			}
		}
	}

	fmt.Println(pares)
}

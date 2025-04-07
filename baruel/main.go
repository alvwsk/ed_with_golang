package main

import (
	"fmt"
	"slices"
)

func main() {
	var fig_totais, fig_baruel int
	fmt.Scan(&fig_totais, &fig_baruel)

	nums_figurinhas := make([]int, fig_baruel)
	for i := range nums_figurinhas {
		fmt.Scan(&nums_figurinhas[i])
	}

	var figs_repetidas []int
	for i := 1; i < fig_baruel; i++ {
		if nums_figurinhas[i] == nums_figurinhas[i-1] {
			figs_repetidas = append(figs_repetidas, nums_figurinhas[i])
		}
	}

	var figs_faltantes []int
	for i := 1; i <= fig_totais; i++ {
		encontrada := slices.Contains(nums_figurinhas, i)
		if !encontrada {
			figs_faltantes = append(figs_faltantes, i)
		}
	}

	if len(figs_repetidas) == 0 {
		fmt.Println("N")
	} else {
		for i, val := range figs_repetidas {
			if i > 0 {
				fmt.Print(" ")
			}
			fmt.Print(val)
		}
		fmt.Println()
	}

	if len(figs_faltantes) == 0 {
		fmt.Println("N")
	} else {
		for i, val := range figs_faltantes {
			if i > 0 {
				fmt.Print(" ")
			}
			fmt.Print(val)
		}
		fmt.Println()
	}
}

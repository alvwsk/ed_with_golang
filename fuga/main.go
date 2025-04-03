package main

import "fmt"

func main() {
	var h, p, f, d int8
	fmt.Scan(&h, &p, &f, &d)

	for {
		if f == p {
			fmt.Println("N")
			break
		}
		if f == h {
			fmt.Println("S")
			break
		}

		f = (f + d + 16) % 16
	}
}

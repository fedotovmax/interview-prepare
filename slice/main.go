package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4, 5}

	for i := 0; i < len(a); i++ {
		fmt.Printf("Value: {%d}; Index: {%d}\n", a[i], i)
	}

	// Происходит копирование слайса внутри range
	for idx, value := range a {
		fmt.Printf("Value: {%d}; Index: {%d}\n", value, idx)
	}

}

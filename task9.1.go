package main

import (
	"fmt"
	"math"
)

func main() {

	fmt.Println("Переполнение.")

	var (
		n uint16
		m uint8
		a int
		b int
	)

	for i := 0; i < math.MaxUint32; i++ {

		n++
		m++

		if m == 0 {
			a++
		}
		if n == 0 {
			b++
		}
	}
	fmt.Println("Переполнений чисел типа uint8, uint16:", a, b)
}

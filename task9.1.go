package main

import (
	"fmt"
)

func main() {

	fmt.Println("Переполнение.")

	const (
		MaxUint32 = 1<<32 - 1
		MaxUint16 = 1<<16 - 1
		MaxUint8  = 1<<8 - 1
	)

	var (
		n int
		m int
	)

	for i := 1; i < MaxUint32; i++ {

		if i%MaxUint16 == 0 {
			n++
		}
		if i%MaxUint8 == 0 {
			m++
		}
	}
	fmt.Println("Переполнений чисел типа uint8, uint16:", n, m)
	fmt.Println(MaxUint32/MaxUint16, MaxUint32/MaxUint8) //проверка
}

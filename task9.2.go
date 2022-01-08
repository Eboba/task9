package main

import (
	"fmt"
	"unsafe"
)

func main() {

	fmt.Println("Минимальный тип данных.")

	const (
		MaxInt8   = 1<<7 - 1
		MinInt8   = -1 << 7
		MaxInt16  = 1<<15 - 1
		MinInt16  = -1 << 15
		MaxInt32  = 1<<31 - 1
		MinInt32  = -1 << 31
		MaxInt64  = 1<<63 - 1
		MinInt64  = -1 << 63
		MaxUint8  = 1<<8 - 1
		MaxUint16 = 1<<16 - 1
		MaxUint32 = 1<<32 - 1
	)

	var (
		b uint8
		c uint16
		d uint32
		e int8
		f int16
		g int32
		h int64
	)

	fmt.Println("Введите первое число (int16):")
	var n int16
	fmt.Scan(&n)

	fmt.Println("Введите второе число (int16):")
	var m int16
	fmt.Scan(&m)

	a := int32(n) * int32(m)

	switch {
	case a >= 0 && a < MaxUint8:
		b = uint8(a)
		fmt.Printf("Тип: %T\nРазмер: %d байт\nЗначение: %v\n", b, unsafe.Sizeof(b), b)

	case a >= 0 && a < MaxUint16:
		c = uint16(a)
		fmt.Printf("Тип: %T\nРазмер: %d байт\nЗначение: %v\n", c, unsafe.Sizeof(c), c)

	case a >= 0 && uint32(a) < MaxUint32:
		d = uint32(a)
		fmt.Printf("Тип: %T\nРазмер: %d байт\nЗначение: %v\n", d, unsafe.Sizeof(d), d)

	case a >= MinInt8:
		e = int8(a)
		fmt.Printf("Тип: %T\nРазмер: %d байт\nЗначение: %v\n", e, unsafe.Sizeof(e), e)

	case a >= MinInt16:
		f = int16(a)
		fmt.Printf("Тип: %T\nРазмер: %d байт\nЗначение: %v\n", f, unsafe.Sizeof(f), f)

	case a >= MinInt32:
		g = int32(a)
		fmt.Printf("Тип: %T\nРазмер: %d байт\nЗначение: %v\n", g, unsafe.Sizeof(g), g)

	default:
		h = int64(a)
		fmt.Printf("Тип: %T\nРазмер: %d байт\nЗначение: %v\n", h, unsafe.Sizeof(h), h)
	}
}

package main

import (
	"fmt"
)

func kurang(a int, b int) {
	c := a - b
	fmt.Println("Hasil Pengurangan Antara", a, "-", b, "Adalah", c)
}
func tambah(a int, b int) int {
	c := a + b
	return c
}

func kalkulator() {
	a := 20
	b := 15
	hasil := tambah(a, b)

	fmt.Println("Hasil Penjumlahan Antara", a, "+", b, "Adalah", hasil)
	kurang(a, b)
}

func main() {
	Belajar1()
	kalkulator()
	Myarray()
	Myassets()
	Servergo()
}

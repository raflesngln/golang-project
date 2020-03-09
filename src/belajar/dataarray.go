package main

import "fmt"

/* import package fmt for library I/O */

func Myarray() {
	// var nama[...]string{}
	var nama [5]string
	nama[0] = "Budi"
	nama[1] = "Mawar"
	nama[2] = "RAfles"
	nama[3] = "Melati"
	nama[4] = "Agus"

	for i := 0; i < len(nama); i++ {
		fmt.Println("Index ke ", i, "=>", nama[i])
	}

}
func arr_range() {
	var fruits = [4]string{"apple", "grape", "banana", "melon"}
	for i, fruit := range fruits {
		fmt.Printf("elemen %d : %s\n", i, fruit)
	}
}
func arr_underscore() {
	var fruits = [4]string{"apple", "grape", "banana", "melon"}
	for _, fruit := range fruits {
		fmt.Printf("nama buah : %s\n", fruit)
	}
}
func arr_multidimensi() {
	var numbers1 = [2][3]int{[3]int{3, 2, 3}, [3]int{3, 4, 5}}
	var numbers2 = [2][3]int{{3, 2, 3}, {3, 4, 5}}

	fmt.Println("numbers1", numbers1) // [[3 2 3] [3 4 5]]
	fmt.Println("numbers2", numbers2) // [[3 2 3] [3 4 5]]

}

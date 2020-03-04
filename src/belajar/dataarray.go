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
	for i, v := range nama {
		fmt.Println("Index ke ", i, "=>", v)
	}

	fmt.Println("=========ATAU===========")

	for index := 0; index < len(nama); index++ {
		fmt.Println("Index ke ", index, "=>", nama[index])
	}

}

package main

import (
	"fmt"
)

var phone="909090909"

const(
	API_URL="http:localhost:8000/api/v1/"
	API_TOKEN="46545645656vtfg"
) 


func main(){
	fmt.Println("Hello world")
	myVar()
	hello := myFunction("John Doe")
	fmt.Println(hello)
}

// variable
func myVar(){
	nama:="rafles"  // hanya di dalam fungsi
	var alamat string="Jakarta" // bisa di inistialisasi ulang
	var phone="0886473247"  // diluar dan dalam fungsi

	var a, b, c int = 1,2,3
	var(
		firstname string = "rafles"
		lastname="nainggolan"
		hobby="Reading"
	)

	fmt.Println("variable",nama," ALamat",alamat,"phone",phone)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(firstname,lastname,hobby,API_URL)
}

func myFunction(name string) string {
	return "Hello, " + name
}
package main

/* import package fmt for library I/O */
import "fmt"

func Belajar1() {
	var firstName string = "Rafles"
	var lastName string
	lastName = "Nainggolan"

	// /*Variabel Tanpa Tipe Data *
	hobbies := "Programmer,"
	hobbies += " browsing,"
	hobbies += " Coding"
	/*Deklarasi Multi Variabel */
	nama, alamat, pekerjaan := "rafles", "jakarta barat", "Programmers"

	fmt.Println("hello world, Keep Learn GO !! \n", firstName, lastName+" My Hobby is ", hobbies)
	fmt.Println("Biodataku adalah :", nama, alamat, pekerjaan)

	/*dengan keyword NEW */
	name := new(string)
	fmt.Println(name)  // 0x20818a220
	fmt.Println(*name) // ""

	// TIPE DATA NUMERIK non-DECIMAL
	var nilai = (((2+6)%3)*4 - 2) / 3
	var positiveNumber uint8 = 89
	var negativeNumber = -1243423644
	fmt.Println("bilangan positif: ", positiveNumber, "isi perhitungan", nilai)
	fmt.Printf("bilangan negatif: %d\n", negativeNumber)

	// TIPE DATA NUMERIK non-DECIMAL
	var decimalNumber = 2.62
	fmt.Printf("bilangan desimal: %f\n", decimalNumber)
	fmt.Printf("bilangan desimal: %.3f\n", decimalNumber)

	// LOGIKA PEMROGRAMAN
	var point = 8
	if point == 10 {
		fmt.Println("lulus dengan nilai sempurna")
	} else if point > 5 {
		fmt.Println("lulus")
	} else if point == 4 {
		fmt.Println("hampir lulus")
	} else {
		fmt.Printf("tidak lulus. nilai anda %d\n", point)
	}
	// SWITCH CASE
	var myvalue = 8
	switch myvalue {
	case 8:
		fmt.Println("perfect")
	case 7:
		fmt.Println("awesome")
	default:
		fmt.Println("not bad")
	}

}

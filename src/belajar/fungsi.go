package main

import (
	"fmt"
	"math"
	"strings"
)

func printMessage(message string, arr []string) {
	var nameString = strings.Join(arr, " ")
	fmt.Println(message, nameString)
}
func akses_fungsi() {
	var names = []string{"John", "Wick"}
	printMessage("halo", names)
}
func funsgsi_return(min, max int) int {
	var value = min + max
	return value
}
func akses_funsgsi_return() {
	var randomValue int
	randomValue = funsgsi_return(2, 10)
	fmt.Println("random number:", randomValue)
}
func multi_return(d float64) (float64, float64) {
	// hitung luas
	var area = math.Pi * math.Pow(d/2, 2)
	// hitung keliling
	var circumference = math.Pi * d
	// kembalikan 2 nilai
	return area, circumference
}
func akses_multi_return() {
	var diameter float64 = 15
	var area, circumference = multi_return(diameter)

	fmt.Printf("luas lingkaran\t\t: %.2f \n", area)
	fmt.Printf("keliling lingkaran\t: %.2f \n", circumference)

}
func f_variadic(numbers ...int) float64 {
	var total int = 0
	for _, number := range numbers {
		total += number
	}

	var avg = float64(total) / float64(len(numbers))
	return avg
}
func akses_f_variadic() {
	/*atau
	var numbers = []int{2, 4, 3, 5, 4, 3, 3, 5, 5, 3}
	var numbers = []string{"sleeping", "eating"}
	var avg = calculate(numbers...)
	*/
	var avg = f_variadic(2, 4, 3, 5, 4, 3, 3, 5, 5, 3)
	var msg = fmt.Sprintf("Rata-rata : %.2f", avg)

	fmt.Println(msg)
}

type FilterCallback func(string) bool

// func f_closure(data []string, callback func(string) bool) []string {
func f_closure(data []string, callback FilterCallback) []string {
	var result []string
	for _, each := range data {
		if filtered := callback(each); filtered {
			result = append(result, each)
		}
	}
	return result
}
func akses_closure() {
	var data = []string{"wick", "jason", "ethan"}
	var dataContainsO = f_closure(data, func(each string) bool {
		return strings.Contains(each, "o") //bernilai true, jika dalam variable each terdapat huruf o
	})
	var dataLenght5 = f_closure(data, func(each string) bool {
		return len(each) == 5 //bernilai true, jika dalam variable each panjangnya 5
	})

	fmt.Println("data asli \t\t:", data)
	// data asli : [wick jason ethan]

	fmt.Println("filter ada huruf \"o\"\t:", dataContainsO)
	// filter ada huruf "o" : [jason]

	fmt.Println("filter jumlah huruf \"5\"\t:", dataLenght5)
	// filter jumlah huruf "5" : [jason ethan]
}
func f_pointer() {
	var numberA int = 4
	var numberB *int = &numberA

	fmt.Println("numberA (value)   :", numberA)  // 4
	fmt.Println("numberA (address) :", &numberA) // 0xc20800a220  memiliki alamat memori yg sama bahkan jika nilai numbA berubah

	fmt.Println("numberB (value)   :", *numberB) // 4
	fmt.Println("numberB (address) :", numberB)  // 0xc20800a220   memiliki alamat memori yg sama bahkan jika nilai numbA berubah

}
func change_pointer(original *int, value int) {
	*original = value
}
func skses_pointer() {
	var number = 4
	fmt.Println("before :", number) // 4

	change(&number, 10)
	fmt.Println("after  :", number) // 10
}
func main() {
	akses_fungsi()
	akses_funsgsi_return()
	akses_f_variadic()
	akses_closure()

}

package main

import (
	"fmt"
)

func f_slice() {
	var fruits = []string{"apple", "grape", "banana", "melon"}
	var newFruits = fruits[0:2]
	// fmt.Println(newFruits) // ["apple", "grape"] get from index 0 till index 2
	fmt.Println(newFruits)

}
func f_cad() {
	var fruits = []string{"apple", "grape", "banana", "melon"}
	fmt.Println(len(fruits)) // len: 4
	fmt.Println(cap(fruits)) // cap: 4
	var aFruits = fruits[0:3]
	fmt.Println(len(aFruits)) // len: 3
	fmt.Println(cap(aFruits)) // cap: 4

}
func f_append() {
	var fruits = []string{"apple", "grape", "banana"}
	var cFruits = append(fruits, "papaya")
	fmt.Println(fruits)  // ["apple", "grape", "banana"]
	fmt.Println(cFruits) // ["apple", "grape", "banana", "papaya"]
}
func f_copy() {
	dst := make([]string, 3)
	src := []string{"watermelon", "pinnaple", "apple", "orange"}
	n := copy(dst, src)
	fmt.Println(dst) // watermelon pinnaple apple
	fmt.Println(src) // watermelon pinnaple apple orange
	fmt.Println(n)   // 3
}
func f_map() {
	var bulan map[string]int
	bulan = map[string]int{} // map bulan dengan key string dan value integer
	var chicken1 = map[string]int{
		"januari":  50,
		"februari": 40,
	}
	bulan["januari"] = 50
	bulan["februari"] = 40
	fmt.Println("januari", bulan["januari"]) // januari 50
	fmt.Println("mei", bulan["mei"])         // mei 0
	fmt.Println("data", chicken1["januari"]) // mei 0
}
func f_map_looping() {
	var chicken = map[string]int{
		"januari":  50,
		"februari": 40,
		"maret":    34,
		"april":    67,
	}
	for key, val := range chicken {
		fmt.Println(key, "  \t:", val)

	}
}
func f_map_delete() {
	var chicken = map[string]int{"januari": 50, "februari": 40}
	fmt.Println(len(chicken)) // 2
	fmt.Println(chicken)
	delete(chicken, "januari")
	fmt.Println(len(chicken)) // 1
	fmt.Println(chicken)
}
func f_map_isexist() {
	var chicken = map[string]int{"januari": 50, "februari": 40}
	var value, isExist = chicken["mei"]
	if isExist {
		fmt.Println(value)
	} else {
		fmt.Println("item is not exists")
	}
}
func slice_and_map() {
	var chickens = []map[string]string{
		map[string]string{"name": "chicken blue", "gender": "male"},
		map[string]string{"name": "chicken red", "gender": "male"},
		map[string]string{"name": "chicken yellow", "gender": "female"},
	}
	/* tipe go terbaru, tiap elemen bisa saja memiliki key yang berbeda-beda
	var chicken2 = []map[string]string{
		{"name": "chicken blue", "gender": "male"},
		{"name": "chicken red", "gender": "male"},
		{"name": "chicken yellow", "gender": "female","alamat":"jakarta"},
	} */
	for _, chicken := range chickens {
		fmt.Println(chicken["name"], chicken["gender"])
	}
}
func main() {
	f_slice()
	f_cad()
	f_append()
	f_copy()
	f_map()
	f_map_looping()
	f_map_delete()
	f_map_isexist()
	slice_and_map()
}

-instal golang default di C:
-buat folder project di D/golang
-ubah environment variable windows GOPATH dengan
  pada user variables tambahkan GOPATH dan valuenya ke folder golang
  D:\golang
-pada folder golang buau 3 folder
  -bin
  -pkg
  -src
-src adalah folder tempat project2 kita nanit
-dalam folder src buat folder projek belejar dan buat sebuah file hello.go
-isi file hello.go denagn:

package main
import "fmt"
func main() {

	fmt.Println("Hello world")
}

-run project dengan masuk cmd pada folder belajar dan run:
 go run .     atau go run hello.go

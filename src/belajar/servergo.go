package main

import (
	"fmt"
	"net/http"
)

func about(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("ini halaman about"))
}
func profile(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("ini halaman Profileku"))
}

func Servergo() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "SelamatBelajar GOLANG!")
	})
	http.HandleFunc("/about", about)
	http.HandleFunc("/profile", profile)

	http.ListenAndServe(":8080", nil)
}

package main

import "net/http"

func Myassets() {
	fsHan := http.FileServer(http.Dir("assets/"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fsHan))
}

package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/pkg/browser"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	fmt.Println("Open Browser http://localhost:8080")
	err := browser.OpenURL("http://localhost:8080")
	if err != nil {
		log.Fatal(err)
	}

	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", http.StripPrefix("/", fs))
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

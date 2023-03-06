package main

import (
	"fmt"
	"log"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Print(w, "test endpoint")
}
func main() {
	http.HandleFunc("/", index)
	err := http.ListenAndServeTLS(":443", "server.crt", "server.key", nil)
	if err != nil {
		log.Fatal("ListenAndServeTLS error: ", err)
	}
}

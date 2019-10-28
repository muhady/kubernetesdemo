package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, first GO container listenning on port 53000")
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":53000", nil))
}

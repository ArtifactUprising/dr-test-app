package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	log.Printf("%v %v %v", r.RemoteAddr, r.Method, r.RequestURI)
	fmt.Fprintf(w, "DrivenRole Test App")
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":9090", nil))
}

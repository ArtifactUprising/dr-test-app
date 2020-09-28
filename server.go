package main

import (
	"fmt"
	"os"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	log.Printf("%v %v %v", r.RemoteAddr, r.Method, r.RequestURI)
	fmt.Fprintf(w, "%s - %s", "DrivenRole Test App", os.Getenv("HOSTNAME"))
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":9090", nil))
}

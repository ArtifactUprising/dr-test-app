package main

import (
	"fmt"
	"os"
	"log"
	"net/http"
	"github.com/newrelic/go-agent"
)

func handler(w http.ResponseWriter, r *http.Request) {
	log.Printf("%v %v %v", r.RemoteAddr, r.Method, r.RequestURI)
	fmt.Fprintf(w, "%s - %s", "DrivenRole Test App", os.Getenv("HOSTNAME"))
}


func main() {
	newrelic_license_key, nr_key_exists := os.LookupEnv("NEWRELIC_LICENSE_KEY")

	if nr_key_exists {
		config := newrelic.NewConfig("DrivenRole Test App", newrelic_license_key)
		app, err := newrelic.NewApplication(config)
		if err != nil {
			log.Fatal("unable to reach new relic api ", err)
		} else {
			log.Println("Connected to NewRelic APM")
		}
		http.HandleFunc(newrelic.WrapHandleFunc(app, "/", handler))
	} else {
		http.HandleFunc("/", handler)
	}
	log.Fatal(http.ListenAndServe(":9090", nil))
}

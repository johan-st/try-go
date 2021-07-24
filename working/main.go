package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	l := log.New(os.Stdout, "product-api", log.LstdFlags)
	helloHandler := handlers.newHello(l)
	// http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
	// 	handlers.Hello(log)
	// })

	// http.HandleFunc("/goodbye", func(rw http.ResponseWriter, r *http.Request) {
	// 	log.Printf("Goodbye World")
	// })

	http.ListenAndServe(":9090", nil)
}

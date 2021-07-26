package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
	"github.com/johan-st/try-go/working/handlers"
)

func main() {

	l := log.New(os.Stdout, "product-api ", log.LstdFlags)

	// create handlers
	productsHandler := handlers.NewProducts(l)

	// create mux and register handlers
	sm := mux.NewRouter()
	// sm.Handle("/products", productsHandler)

	getRouter := sm.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/", productsHandler.GetProducts)

	postRouter := sm.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/", productsHandler.AddProduct)

	putRouter := sm.Methods(http.MethodPut).Subrouter()
	putRouter.HandleFunc("/{id:[0-9]+}", productsHandler.UpdateProduct)

	// create and configure server
	s := &http.Server{
		Addr:         ":9090",
		Handler:      sm,
		IdleTimeout:  2 * time.Second,
		ReadTimeout:  500 * time.Millisecond,
		WriteTimeout: 500 * time.Millisecond,
	}

	go func() {
		l.Println("Starting server on port 9090")
		err := s.ListenAndServe()
		if err != nil {
			l.Fatal(err)
		}
	}()

	signalChan := make(chan os.Signal)
	signal.Notify(signalChan, os.Interrupt)
	signal.Notify(signalChan, os.Kill)

	sig := <-signalChan
	l.Println("Graceful shutdown", sig)

	// No context leakage because of shutdown
	tc, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(tc)
}

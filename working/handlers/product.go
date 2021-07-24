package handlers

import (
	"log"
	"net/http"

	"github.com/johan-st/try-go/working/data"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

func (p *Products) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		p.getProducts(rw, r)
		return
	}
	if r.Method == http.MethodPost {
		p.addProduct(rw, r)
		return
	}

	// catch-all
	rw.WriteHeader(http.StatusNotImplemented)
}

func (p *Products) getProducts(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle GET products")
	list := data.GetProducts()
	err := list.WriteJSON(rw)

	if err != nil {
		http.Error(rw, "Unable to send json", http.StatusInternalServerError)
	}

}

func (p *Products) addProduct(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle POST products")

	prod := &data.Product{}
	err := prod.CreateFromJson(r.Body)
	if err != nil {
		http.Error(rw, "Unable to create product from JSON", http.StatusBadRequest)
	}
	data.AddProduct(prod)
	// p.l.Printf("Prod: %#v", prod)
}

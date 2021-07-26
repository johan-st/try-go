package handlers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/johan-st/try-go/working/data"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

func (p *Products) GetProducts(rw http.ResponseWriter, req *http.Request) {
	p.l.Println("Handle GET products")
	list := data.GetProducts()
	err := list.WriteJSON(rw)

	if err != nil {
		http.Error(rw, "Unable to send json", http.StatusInternalServerError)
	}

}

func (p *Products) UpdateProduct(rw http.ResponseWriter, req *http.Request) {
	p.l.Println("Handle PUT product")
	vars := mux.Vars(req)
	id, errId := strconv.Atoi(vars["id"])
	if errId != nil {
		http.Error(rw, "unable to parse id from url", http.StatusBadRequest)
		return
	}
	prod := &data.Product{}
	prod.CreateFromJson(req.Body)
	err := data.UpdateProduct(id, prod)
	if err != nil {
		http.Error(rw, "unable to process request", http.StatusBadRequest)
	}
}

func (p *Products) AddProduct(rw http.ResponseWriter, req *http.Request) {
	p.l.Println("Handle POST products")

	prod := &data.Product{}
	err := prod.CreateFromJson(req.Body)
	if err != nil {
		http.Error(rw, "Unable to create product from JSON", http.StatusBadRequest)
	}
	data.AddProduct(prod)
	// p.l.Printf("Prod: %#v", prod)
}

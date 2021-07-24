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
		p.GetProducts(rw, r)
		return
	}

	// catch-all
	rw.WriteHeader(http.StatusNotImplemented)
}

func (p *Products) GetProducts(rw http.ResponseWriter, r *http.Request) {
	list := data.GetProducts()
	err := list.WriteJSON(rw)

	if err != nil {
		http.Error(rw, "Unable to send json", http.StatusInternalServerError)
	}

}

package handlers

import (
	"golangProjects/Microservice/Gorillamux/data"
	"net/http"
)

// swagger:route GET /products products listProducts
// Returns a list of product
// responses:
//   200: productResponse
// GetProducts returns the products from the data store

func (p *Products) GetProducts(w http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle in GET Product")

	lp := data.GetProducts()

	err := lp.ToJson(w)
	if err != nil {
		http.Error(w, "unable to marshal json", http.StatusBadRequest)
	}
}

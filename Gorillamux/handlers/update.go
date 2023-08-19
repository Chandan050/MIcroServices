package handlers

import (
	"golangProjects/Microservice/Gorillamux/data"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func (p *Products) UpdateProducts(w http.ResponseWriter, r *http.Request) {

	// mux.Vars will extract the value from url at end point r - client request
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	p.l.Println("Handle put products", id)

	if id > len(data.ProductList) {
		p.l.Println("entered id is not exit in database")
	}
	prod := r.Context().Value(KeyProduct{}).(data.Product)

	err := data.UpdateProduct(id, &prod)
	if err == data.ErrProductNotFond {
		http.Error(w, "product not found", http.StatusBadRequest)
		return
	}

	if err != nil {
		http.Error(w, "product not found", http.StatusInternalServerError)
		return
	}

}

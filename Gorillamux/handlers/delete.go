package handlers

import (
	"fmt"
	"golangProjects/Microservice/Gorillamux/data"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

//swagger:route DELETE /products{id} products deleteproducts
//Return a list of products
// responses:
//	201: nocontent
//DeleteProduct delete a product from the database

func (p *Products) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	p.l.Println("Handle delete products")

	dp, err := data.DeleteProduct(id, w, r)

	if err != nil {
		http.Error(w, "product not found", http.StatusInternalServerError)
		return
	}

	p.l.Printf(fmt.Sprintf("%v product deleted ", dp))

}

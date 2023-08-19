//Package Classification of Product-API
//
//Documentation for Product-api
//
// Schemes: http
// BasePath: /
// Version: 1.0.0
//
// Consumes:
//  	-appication/json
//
// Produces:
//   -application/json
//swagger:meta
package handlers

import (
	"context"
	"fmt"
	"golangProjects/Microservice/Gorillamux/data"
	"log"
	"net/http"
)

// A list of products returns in the response
// swagger:response productsResponse
type productsResponseWrapper struct {
	// All products in  the system
	// in: body
	Body []data.Product
}

// swagger:response noContent
type productNoContent struct {
	// Message to show that no content was found on this endpoint
}

// swagger:parameters deleteProduct
type productIDParameter struct {
	// The id of the product to delete from the database
	// in: path
	// required: true
	ID int `json:"id"`
}

// produts is a http.handler
type Products struct {
	l *log.Logger
}

// newproducts creates a productes handler with the given logger
func NewProduct(l *log.Logger) *Products {
	return &Products{l}
}

func (p *Products) AddProducts(w http.ResponseWriter, r *http.Request) {

	p.l.Println("Handle Post products")

	// r.context *containes data* Value is r.context contains value with is hodling is KeyProduct{} that shoud be type of Product which is in data file
	prod := r.Context().Value(KeyProduct{}).(data.Product)
	data.AddProducts(&prod)
}

type KeyProduct struct{}

//MiddlewareProductValidation validate the input

func (p *Products) MiddlewareProductValidation(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		prod := data.Product{}
		err := prod.FromJson(r.Body)

		if err != nil {
			p.l.Println("[ERROR] deserializing product", err)
			http.Error(w, "invalid request body", 400)
			return
		}
		// context.Background is empty interface hete we use r(request).Context as interface
		// {keyproduct : prod} = key-value type
		ctx := context.WithValue(r.Context(), KeyProduct{}, prod)

		// prod validation
		err = prod.Validate()
		if err != nil {
			p.l.Println("[ERROR] validating Product", err)
			http.Error(w, fmt.Sprintf("ERROR while validating Product %v", err), http.StatusBadRequest)
			return
		}

		// stores the context value in variable req
		req := r.WithContext(ctx)

		// next is an hadler we are sending data to w, r who are passing
		next.ServeHTTP(w, req)
	})
}

package data

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"regexp"
	"time"

	"github.com/go-playground/validator"
)

// Product defines the structure for API product
// swagger:model
type Product struct {
	// the id for this user
	//
	// required:true
	//min:1
	ID          int     `json:"id"`
	Name        string  `json:"name" validate:"required"`
	Description string  `json:"description"`
	Price       float64 `json:"price" validate:"gt=0.0"`
	SKU         string  `json:"sku"`
	createdOn   string  `json:"-"`
	updatedOn   string  `json:"-"`
	DeletedOn   string  `json:"-"`
}

func (p *Product) FromJson(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(p)
}

func (p *Product) Validate() error {
	validate := validator.New()
	validate.RegisterValidation("sku", validateSKU)
	return validate.Struct(p)

}

func validateSKU(fl validator.FieldLevel) bool {
	re := regexp.MustCompile(`[a-z]+-[a-z]+-[a-z]+`)
	matches := re.FindAllString(fl.Field().String(), -1)

	if len(matches) != 1 {
		return false
	}
	return true

}

// products is a collection of product
type Products []*Product

// encode will write in to io.Writer dirctly so if we "json.marshal" = it as to convert and write to http so on to io.writer
func (p *Products) ToJson(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}

// GetProducts retunr a list of products
func GetProducts() Products {
	return ProductList
}

func AddProducts(p *Product) {
	p.ID = getNextId()
	ProductList = append(ProductList, p)

}
func UpdateProduct(id int, p *Product) error {
	_, Pos, err := findProduct(id)
	if err != nil {
		return err
	}
	p.ID = id
	ProductList[Pos] = p
	return nil

}
func DeleteProduct(id int, w http.ResponseWriter, r *http.Request) (*Product, error) {
	dp, n, err := findProduct(id)
	if err != nil {
		return nil, err
	}

	ProductList = append(ProductList[:n], ProductList[n+1:]...)
	return dp, nil
}

var ErrProductNotFond = fmt.Errorf("productNotFond")

func findProduct(id int) (*Product, int, error) {
	for i, p := range ProductList {
		if p.ID == id {
			return p, i, nil
		}

	}
	return nil, -1, ErrProductNotFond
}
func getNextId() int {
	pi := ProductList[len(ProductList)-1]
	return pi.ID + 1
}

var ProductList = []*Product{
	&Product{
		ID:          1,
		Name:        "capchino",
		Description: "frothy milky coffee",
		Price:       90,
		SKU:         "abc123",
		createdOn:   time.Now().UTC().String(),
		updatedOn:   time.Now().UTC().String(),
	},
	&Product{
		ID:          2,
		Name:        "salad",
		Description: "forty milky coffee",
		Price:       180,
		SKU:         "bcd121",
		createdOn:   time.Now().UTC().String(),
		updatedOn:   time.Now().UTC().String(),
	},
}

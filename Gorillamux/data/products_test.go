package data

import "testing"

func TestCheckValidation(t *testing.T) {
	// TODO: Implement test

	p := &Product{
		Name:  "nics",
		Price: 3.56,
		SKU:   "abc-abc-abs",
	}

	err := p.Validate()

	if err != nil {
		t.Fatal(err)
	}

}

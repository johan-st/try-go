package data

import "testing"

func TestValidation(t *testing.T) {
	p := &Product{
		Name:  "bobbins",
		Price: 0.01,
		SKU:   "abc-abc-dbcd",
	}

	err := p.Validate()
	if err != nil {
		t.Fatal(err)
	}
}

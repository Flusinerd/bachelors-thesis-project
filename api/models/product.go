package models

import (
	"encoding/json"
	"io"
)

type Product struct {
	Id          string  `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
}

func FromJSON(r io.Reader, v interface{}) error {
	return json.NewDecoder(r).Decode(v)
}

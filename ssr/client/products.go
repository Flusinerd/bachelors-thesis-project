package client

import (
	"api/models"
	"net/http"
)

func GetProducts() []models.Product {
	resp, err := http.Get("http://localhost:8080/api/products")
	if err != nil {
		return nil
	}
	defer resp.Body.Close()

	var products []models.Product
	err = models.FromJSON(resp.Body, &products)
	if err != nil {
		return nil
	}

	return products
}

func GetProduct(id string) *models.Product {
	resp, err := http.Get("http://localhost:8080/api/products/" + id)
	if err != nil {
		return nil
	}

	defer resp.Body.Close()

	var product models.Product
	err = models.FromJSON(resp.Body, &product)
	if err != nil {
		return nil
	}

	return &product
}

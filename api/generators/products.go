package generators

import (
	"api/models"
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/brianvoe/gofakeit/v6"
)

func generateProducts() map[string]models.Product {
	productsMap := map[string]models.Product{}

	for i := 0; i < 500; i++ {
		// Create a product using gofakeit
		product := models.Product{
			Id:          gofakeit.UUID(),
			Name:        gofakeit.ProductName(),
			Description: gofakeit.ProductDescription(),
			Price:       gofakeit.Price(0, 1000),
		}

		productsMap[product.Id] = product
	}

	return productsMap
}

func SaveProductsToJson(products *map[string]models.Product) {
	assetsDir := "assets"
	fileName := "products.json"

	path := filepath.Join(assetsDir, fileName)
	absPath, _ := filepath.Abs(path)

	os.Remove(absPath)

	file, _ := os.Create(absPath)
	defer file.Close()

	json := json.NewEncoder(file)
	json.SetIndent("", "  ")
	json.Encode(products)
}

func LoadProductsFromJson() map[string]models.Product {
	assetsDir := "assets"
	fileName := "products.json"

	path := filepath.Join(assetsDir, fileName)
	absPath, _ := filepath.Abs(path)

	file, err := os.Open(absPath)
	if err != nil {
		if os.IsNotExist(err) {
			products := generateProducts()
			return products
		}
		panic(err)
	}
	defer file.Close()

	var products map[string]models.Product
	json.NewDecoder(file).Decode(&products)

	return products
}

package repositories

import (
	"api/models"
	"sort"
)

type productsRepository struct {
	products map[string]models.Product
}

type ProductsRepository interface {
	GetProducts() []models.Product
	GetProductById(id string) *models.Product
}

func NewProductsRepository(products map[string]models.Product) ProductsRepository {
	return &productsRepository{products}
}

func (repo *productsRepository) GetProducts() []models.Product {
	slice := make([]models.Product, 0, len(repo.products))
	for _, product := range repo.products {
		slice = append(slice, product)
	}

	sort.Slice(slice, func(i, j int) bool {
		return slice[i].Id < slice[j].Id
	})

	return slice
}

func (repo *productsRepository) GetProductById(id string) *models.Product {
	product := repo.products[id]
	return &product
}

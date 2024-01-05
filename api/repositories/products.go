package repositories

import "api/models"

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

	return slice
}

func (repo *productsRepository) GetProductById(id string) *models.Product {
	product := repo.products[id]
	return &product
}

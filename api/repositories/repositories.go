package repositories

import "api/generators"

type Repositories struct {
	Products ProductsRepository
}

func NewRepositories() *Repositories {
	return &Repositories{
		Products: NewProductsRepository(generators.LoadProductsFromJson()),
	}
}

package state

import "api/models"

type Customer struct {
	id   string
	Cart []models.Product
}

type CustomerState struct {
	Customers map[string]Customer
}

func NewCustomerState() *CustomerState {
	return &CustomerState{
		Customers: map[string]Customer{},
	}
}

func NewCustomer(id string) *CustomerState {
	return &CustomerState{
		Customers: map[string]Customer{
			id: {
				id:   id,
				Cart: []models.Product{},
			},
		},
	}
}

func (c *CustomerState) AddProduct(customerId string, product models.Product) {
	c.Customers[customerId] = Customer{
		id:   customerId,
		Cart: append(c.Customers[customerId].Cart, product),
	}
}

func (c *CustomerState) RemoveProduct(id string, product models.Product) {
	c.Customers[id] = Customer{
		id:   id,
		Cart: removeProduct(c.Customers[id].Cart, product),
	}
}

func (c *CustomerState) GetCart(customerId string) []models.Product {
	return c.Customers[customerId].Cart
}

func removeProduct(products []models.Product, product models.Product) []models.Product {
	for i, p := range products {
		if p.Id == product.Id {
			return append(products[:i], products[i+1:]...)
		}
	}
	return products
}

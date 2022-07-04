package repository

import "enigmacamp.com/go_product/model"

type ProductRepository interface {
	Add(newProduct *model.Product) error
	Retrieve() []model.Product
}

type productRepository struct {
	productDb []model.Product
}

func (p *productRepository) Add(newProduct *model.Product) error {
	p.productDb = append(p.productDb, *newProduct)
	return nil
}

func (p *productRepository) Retrieve() []model.Product {
	return nil
}

func NewProductRepository() ProductRepository {
	repo := new(productRepository)
	return repo
}

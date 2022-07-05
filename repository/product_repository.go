package repository

import (
	"enigmacamp.com/go_product/model"
	"gorm.io/gorm"
)

type ProductRepository interface {
	Add(newProduct *model.Product) error
	Retrieve() []model.Product
}

type productRepository struct {
	db *gorm.DB
}

func (p *productRepository) Add(newProduct *model.Product) error {
	return p.db.Create(newProduct).Error
}

func (p *productRepository) Retrieve() []model.Product {
	var products []model.Product
	p.db.Find(&products)
	return products
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	repo := new(productRepository)
	repo.db = db
	return repo
}

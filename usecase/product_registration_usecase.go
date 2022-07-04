package usecase

import (
	"enigmacamp.com/go_product/model"
	"enigmacamp.com/go_product/repository"
)

type ProductRegistrationUseCase interface {
	Register(newProduct *model.Product) error
}

type productRegistrationUseCase struct {
	repo repository.ProductRepository
}

func (a *productRegistrationUseCase) Register(newProduct *model.Product) error {
	return a.repo.Add(newProduct)
}

func NewProductRegistrationUseCase(repo repository.ProductRepository) ProductRegistrationUseCase {
	return &productRegistrationUseCase{
		repo: repo,
	}
}

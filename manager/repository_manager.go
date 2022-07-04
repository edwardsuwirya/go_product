package manager

import "enigmacamp.com/go_product/repository"

type RepositoryManager interface {
	ProductRepo() repository.ProductRepository
}

type repositoryManager struct {
}

func (r *repositoryManager) ProductRepo() repository.ProductRepository {
	return repository.NewProductRepository()
}

func NewRepositoryManager() RepositoryManager {
	return &repositoryManager{}
}

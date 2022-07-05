package manager

import "enigmacamp.com/go_product/repository"

type RepositoryManager interface {
	ProductRepo() repository.ProductRepository
}

type repositoryManager struct {
	infraManager InfraManager
}

func (r repositoryManager) ProductRepo() repository.ProductRepository {
	return repository.NewProductRepository(r.infraManager.DbConn())
}

func NewRepositoryManager(manager InfraManager) RepositoryManager {
	return repositoryManager{
		infraManager: manager,
	}
}

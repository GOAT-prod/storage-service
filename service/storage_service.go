package service

import (
	"fmt"
	"storage-service/database"
	"storage-service/domain"
	"storage-service/domain/mappings"
	"storage-service/tools/storagecontext"
)

type StorageService interface {
	GetProducts(ctx storagecontext.StorageContext, limit int, page int) ([]domain.Product, error)
	SaveProduct(ctx storagecontext.StorageContext, product domain.Product) error
	RemoveProduct(ctx storagecontext.StorageContext, productId int) error
	UpdateProduct(ctx storagecontext.StorageContext, product domain.Product) error
}

type StorageServiceImpl struct {
	repo database.StorageRepository
}

func NewStorageService(repos database.StorageRepository) StorageService {
	return &StorageServiceImpl{
		repo: repos,
	}
}

func (s *StorageServiceImpl) GetProducts(ctx storagecontext.StorageContext, limit int, page int) ([]domain.Product, error) {
	dbProducts, err := s.repo.GetProducts(ctx, limit, page*limit)
	if err != nil {
		return nil, fmt.Errorf("error getting products: %w", err)
	}

	products := make([]domain.Product, 0, len(dbProducts))

	for i := range dbProducts {
		dbBrand, bErr := s.repo.GetBrands(ctx, []int{dbProducts[i].BrandId})
		if bErr != nil {
			ctx.Log().Error(fmt.Sprintf("error getting product [id %d] brand: %s", dbProducts[i].BrandId, err))
			continue
		}

		dbFactory, fErr := s.repo.GetFactories(ctx, []int{dbProducts[i].FactoryId})
		if fErr != nil {
			ctx.Log().Error(fmt.Sprintf("error getting product [id %d] factory: %s", dbProducts[i].FactoryId, err))
			continue
		}

		products = append(products, mappings.ToDomainProduct(dbProducts[i], dbBrand[0], dbFactory[0]))
	}

	return products, nil
}

func (s *StorageServiceImpl) SaveProduct(ctx storagecontext.StorageContext, product domain.Product) error {
	return s.repo.AddProduct(ctx, mappings.ToDbProduct(product))
}

func (s *StorageServiceImpl) RemoveProduct(ctx storagecontext.StorageContext, productId int) error {
	return s.repo.DeleteProduct(ctx, productId)
}

func (s *StorageServiceImpl) UpdateProduct(ctx storagecontext.StorageContext, product domain.Product) error {
	return s.repo.UpdateProduct(ctx, mappings.ToDbProduct(product))
}

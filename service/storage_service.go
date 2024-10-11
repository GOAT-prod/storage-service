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
	dbProducts, err := s.repo.GetAllProducts(ctx, limit, page*limit)
	if err != nil {
		return nil, fmt.Errorf("error getting products: %w", err)
	}

	products := make([]domain.Product, 0, len(dbProducts))

	for _, dbProduct := range dbProducts {
		fullDbProduct := dbProduct

		fullDbProduct.Items, err = s.repo.GetProductItems(ctx, []int{dbProduct.Id})
		if err != nil {
			ctx.Log().Error(fmt.Sprintf("error getting product [id %d] items: %s", fullDbProduct.Id, err))
			continue
		}

		fullDbProduct.Images, err = s.repo.GetImages(ctx, []int{dbProduct.Id})
		if err != nil {
			ctx.Log().Error(fmt.Sprintf("error getting product [id %d] images: %s", fullDbProduct.Id, err))
			continue
		}

		fullDbProduct.Materials, err = s.repo.GetMaterials(ctx, []int{dbProduct.Id})
		if err != nil {
			ctx.Log().Error(fmt.Sprintf("error getting product [id %d] materials: %s", fullDbProduct.Id, err))
			continue
		}

		dbBrand, bErr := s.repo.GetBrands(ctx, []int{dbProduct.BrandId})
		if bErr != nil {
			ctx.Log().Error(fmt.Sprintf("error getting product [id %d] brand: %s", fullDbProduct.Id, err))
			continue
		}

		dbFactory, fErr := s.repo.GetFactories(ctx, []int{dbProduct.FactoryId})
		if fErr != nil {
			ctx.Log().Error(fmt.Sprintf("error getting product [id %d] factory: %s", fullDbProduct.Id, err))
			continue
		}

		products = append(products, mappings.ToDomainProduct(fullDbProduct, dbBrand[0], dbFactory[0]))
	}

	return products, nil
}

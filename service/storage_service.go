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
	GetProduct(ctx storagecontext.StorageContext, productId int) (domain.Product, error)
}

type StorageServiceImpl struct {
	storageRepository database.StorageRepository
	logsRepository    database.LogsRepository
}

func NewStorageService(sRepos database.StorageRepository, lRepos database.LogsRepository) StorageService {
	return &StorageServiceImpl{
		storageRepository: sRepos,
		logsRepository:    lRepos,
	}
}

func (s *StorageServiceImpl) GetProduct(ctx storagecontext.StorageContext, productId int) (domain.Product, error) {
	defer func() {
		getProductParams := struct{ ProductId int }{ProductId: productId}
		if err := s.logsRepository.Log(ctx, "get product", getProductParams); err != nil {
			ctx.Log().Error(fmt.Sprintf("не удалось записать лог: %v", err))
		}
	}()

	dbProduct, err := s.storageRepository.GetProduct(ctx, productId)
	if err != nil {
		return domain.Product{}, err
	}

	dbBrand, err := s.storageRepository.GetBrands(ctx, []int{dbProduct.BrandId})
	if err != nil {
		return domain.Product{}, err
	}

	dbFactory, err := s.storageRepository.GetFactories(ctx, []int{dbProduct.FactoryId})
	if err != nil {
		return domain.Product{}, err
	}

	return mappings.ToDomainProduct(dbProduct, dbBrand[0], dbFactory[0]), nil
}

func (s *StorageServiceImpl) GetProducts(ctx storagecontext.StorageContext, limit int, page int) ([]domain.Product, error) {
	defer func() {
		getProductsParams := struct {
			Limit int
			Page  int
		}{Limit: limit, Page: page}
		if err := s.logsRepository.Log(ctx, "get products", getProductsParams); err != nil {
			ctx.Log().Error(fmt.Sprintf("не удалось записать лог: %v", err))
		}
	}()

	dbProducts, err := s.storageRepository.GetProducts(ctx, limit, page*limit)
	if err != nil {
		return nil, fmt.Errorf("error getting products: %w", err)
	}

	products := make([]domain.Product, 0, len(dbProducts))

	for i := range dbProducts {
		dbBrand, bErr := s.storageRepository.GetBrands(ctx, []int{dbProducts[i].BrandId})
		if bErr != nil {
			ctx.Log().Error(fmt.Sprintf("error getting product [id %d] brand: %s", dbProducts[i].BrandId, err))
			continue
		}

		dbFactory, fErr := s.storageRepository.GetFactories(ctx, []int{dbProducts[i].FactoryId})
		if fErr != nil {
			ctx.Log().Error(fmt.Sprintf("error getting product [id %d] factory: %s", dbProducts[i].FactoryId, err))
			continue
		}

		products = append(products, mappings.ToDomainProduct(dbProducts[i], dbBrand[0], dbFactory[0]))
	}

	return products, nil
}

func (s *StorageServiceImpl) SaveProduct(ctx storagecontext.StorageContext, product domain.Product) error {
	defer func() {
		if err := s.logsRepository.Log(ctx, "save product", product); err != nil {
			ctx.Log().Error(fmt.Sprintf("не удалось записать лог: %v", err))
		}
	}()

	return s.storageRepository.AddProduct(ctx, mappings.ToDbProduct(product))
}

func (s *StorageServiceImpl) RemoveProduct(ctx storagecontext.StorageContext, productId int) error {
	defer func() {
		deleteParams := struct{ ProductId int }{ProductId: productId}
		if err := s.logsRepository.Log(ctx, "delete product", deleteParams); err != nil {
			ctx.Log().Error(fmt.Sprintf("не удалось записать лог: %v", err))
		}
	}()

	return s.storageRepository.DeleteProduct(ctx, productId)
}

func (s *StorageServiceImpl) UpdateProduct(ctx storagecontext.StorageContext, product domain.Product) error {
	defer func() {
		if err := s.logsRepository.Log(ctx, "update product", product); err != nil {
			ctx.Log().Error(fmt.Sprintf("не удалось записать лог: %v", err))
		}
	}()

	return s.storageRepository.UpdateProduct(ctx, mappings.ToDbProduct(product))
}

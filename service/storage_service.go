package service

import (
	"github.com/samber/lo"
	"storage-service/database"
	"storage-service/domain"
	"storage-service/domain/mappings"
	"storage-service/tools/storagecontext"
)

type StorageService interface {
	GetProducts(ctx storagecontext.StorageContext, limit int) ([]domain.ProductInfo, error)
}

type StorageServiceImpl struct {
	repo database.StorageRepository
}

func NewStorageService(repos database.StorageRepository) StorageService {
	return &StorageServiceImpl{
		repo: repos,
	}
}

func (s *StorageServiceImpl) GetProducts(ctx storagecontext.StorageContext, limit int) ([]domain.ProductInfo, error) {
	var (
		products []database.DbProductInfo
		err      error
	)

	switch limit {
	case 0:
		products, err = s.repo.GetAllProducts(ctx)
	default:
		products, err = s.repo.GetLimitProducts(ctx, limit)
	}

	if err != nil {
		return nil, err
	}

	productIds := lo.Map(products, func(item database.DbProductInfo, _ int) int {
		return item.Id
	})

	images, err := s.repo.GetImages(ctx, productIds)
	if err != nil {
		return nil, err
	}

	materials, err := s.repo.GetMaterials(ctx, productIds)
	if err != nil {
		return nil, err
	}

	result := mappings.ToProductInfos(products)

	for i := range result {
		result[i].Images = lo.Map(lo.Filter(images, func(item database.DbImage, _ int) bool { return item.ProductId == result[i].Id }), func(item database.DbImage, _ int) string { return item.Url })
		result[i].Materials = lo.Map(lo.Filter(materials, func(item database.DbMaterial, _ int) bool { return item.ProductId == result[i].Id }), func(item database.DbMaterial, _ int) string { return item.Name })
	}

	return result, nil
}

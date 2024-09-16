package mappings

import (
	"github.com/samber/lo"
	"storage-service/database"
	"storage-service/domain"
)

func ToDbProductInfo(info domain.ProductInfo) database.DbProductInfo {
	return database.DbProductInfo{
		Id:           info.Id,
		Name:         info.Name,
		Price:        info.Price,
		Discount:     info.Discount,
		Size:         info.Size,
		Color:        info.Color,
		TypeId:       info.Type.Id,
		TypeName:     info.Type.Name,
		CategoryId:   info.Category.Id,
		CategoryName: info.Category.Name,
	}
}

func ToProductInfo(info database.DbProductInfo) domain.ProductInfo {
	return domain.ProductInfo{
		Id:       info.Id,
		Name:     info.Name,
		Price:    info.Price,
		Discount: info.Discount,
		Size:     info.Size,
		Color:    info.Color,
		Type:     ToClothesType(info.TypeId, info.TypeName),
		Category: ToClothesCategory(info.CategoryId, info.CategoryName),
	}
}

func ToClothesType(id int, name string) domain.ClothesType {
	return domain.ClothesType{
		Id:   id,
		Name: name,
	}
}

func ToClothesCategory(id int, name string) domain.ClothesCategory {
	return domain.ClothesCategory{
		Id:   id,
		Name: name,
	}
}

func ToDbProductInfos(info []domain.ProductInfo) []database.DbProductInfo {
	return lo.Map(info, func(item domain.ProductInfo, _ int) database.DbProductInfo {
		return ToDbProductInfo(item)
	})
}

func ToProductInfos(info []database.DbProductInfo) []domain.ProductInfo {
	return lo.Map(info, func(item database.DbProductInfo, _ int) domain.ProductInfo {
		return ToProductInfo(item)
	})
}

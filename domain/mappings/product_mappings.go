package mappings

import (
	"github.com/samber/lo"
	"storage-service/database"
	"storage-service/domain"
)

func ToDbProduct(info domain.Product) database.Product {
	return database.Product{
		Id:          info.Id,
		BrandId:     info.Brand.Id,
		FactoryId:   info.Factory.Id,
		Name:        info.Name,
		Description: info.Description,
		Price:       info.Price,
		Items:       ToDbProductItems(info.Items, info.Id),
		Materials:   ToDbProductMaterials(info.Materials, info.Id),
		Images:      ToDbProductImages(info.Images, info.Id),
	}
}

func ToDomainProduct(info database.Product, brand database.Brand, factory database.Factory) domain.Product {
	return domain.Product{
		Id:          info.Id,
		Brand:       ToDomainBrand(brand),
		Factory:     ToDomainFactory(factory),
		Name:        info.Name,
		Description: info.Description,
		Price:       info.Price,
		Items:       ToDomainProductItems(info.Items),
		Materials:   ToDomainProductMaterials(info.Materials),
		Images:      ToDomainProductImages(info.Images),
	}
}

func ToDbProductItem(item domain.ProductItem, productId int) database.ProductItem {
	return database.ProductItem{
		Id:         item.Id,
		ProductId:  productId,
		StockCount: item.StockCount,
		Size:       item.Size,
		Weight:     item.Weight,
		Color:      item.Color,
	}
}

func ToDomainProductItem(item database.ProductItem) domain.ProductItem {
	return domain.ProductItem{
		Id:         item.Id,
		StockCount: item.StockCount,
		Size:       item.Size,
		Weight:     item.Weight,
		Color:      item.Color,
	}
}

func ToDbProductItems(items []domain.ProductItem, productId int) []database.ProductItem {
	return lo.Map(items, func(item domain.ProductItem, _ int) database.ProductItem {
		return ToDbProductItem(item, productId)
	})
}

func ToDomainProductItems(items []database.ProductItem) []domain.ProductItem {
	return lo.Map(items, func(item database.ProductItem, _ int) domain.ProductItem {
		return ToDomainProductItem(item)
	})
}

func ToDbProductMaterial(material domain.ProductMaterial, productId int) database.ProductMaterial {
	return database.ProductMaterial{
		Id:        material.Id,
		ProductId: productId,
		Name:      material.Name,
	}
}

func ToDomainProductMaterial(material database.ProductMaterial) domain.ProductMaterial {
	return domain.ProductMaterial{
		Id:   material.Id,
		Name: material.Name,
	}
}

func ToDbProductMaterials(materials []domain.ProductMaterial, productId int) []database.ProductMaterial {
	return lo.Map(materials, func(item domain.ProductMaterial, _ int) database.ProductMaterial {
		return ToDbProductMaterial(item, productId)
	})
}

func ToDomainProductMaterials(materials []database.ProductMaterial) []domain.ProductMaterial {
	return lo.Map(materials, func(item database.ProductMaterial, _ int) domain.ProductMaterial {
		return ToDomainProductMaterial(item)
	})
}

func ToDbProductImage(image domain.ProductImage, productId int) database.ProductImage {
	return database.ProductImage{
		Id:        image.Id,
		ProductId: productId,
		ImageUrl:  image.ImageUrl,
	}
}

func ToDomainProductImage(image database.ProductImage) domain.ProductImage {
	return domain.ProductImage{
		Id:       image.Id,
		ImageUrl: image.ImageUrl,
	}
}

func ToDbProductImages(images []domain.ProductImage, productId int) []database.ProductImage {
	return lo.Map(images, func(item domain.ProductImage, _ int) database.ProductImage {
		return ToDbProductImage(item, productId)
	})
}

func ToDomainProductImages(images []database.ProductImage) []domain.ProductImage {
	return lo.Map(images, func(item database.ProductImage, _ int) domain.ProductImage {
		return ToDomainProductImage(item)
	})
}

func ToDbBrand(brand domain.Brand) database.Brand {
	return database.Brand{
		Id:   brand.Id,
		Name: brand.Name,
	}
}

func ToDomainBrand(brand database.Brand) domain.Brand {
	return domain.Brand{
		Id:   brand.Id,
		Name: brand.Name,
	}
}

func ToDbFactory(factory domain.Factory) database.Factory {
	return database.Factory{
		Id:      factory.Id,
		Name:    factory.Name,
		Country: factory.Country,
		City:    factory.City,
		Address: factory.Address,
	}
}

func ToDomainFactory(factory database.Factory) domain.Factory {
	return domain.Factory{
		Id:      factory.Id,
		Name:    factory.Name,
		Country: factory.Country,
		City:    factory.City,
		Address: factory.Address,
	}
}

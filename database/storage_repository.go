package database

import (
	"context"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	"storage-service/database/sql"
	"storage-service/tools/storagecontext"
)

type StorageRepository interface {
	GetProducts(ctx storagecontext.StorageContext, limit int, offset int) ([]Product, error)
	GetFactories(ctx storagecontext.StorageContext, factoryIds []int) (factories []Factory, err error)
	GetBrands(ctx storagecontext.StorageContext, brandIds []int) (brands []Brand, err error)
	AddProduct(ctx storagecontext.StorageContext, product Product) error
	UpdateProduct(ctx storagecontext.StorageContext, product Product) error
	DeleteProduct(ctx storagecontext.StorageContext, productId int) error
	GetProduct(ctx storagecontext.StorageContext, productId int) (product Product, err error)

	InsertForTest() error
}

type StorageRepositoryImpl struct {
	db *sqlx.DB
}

func NewStorageRepository(db *sqlx.DB) StorageRepository {
	return &StorageRepositoryImpl{
		db: db,
	}
}

func (r *StorageRepositoryImpl) GetProduct(ctx storagecontext.StorageContext, productId int) (product Product, err error) {
	if err = r.db.GetContext(ctx.Ctx(), &product, sql.GetProductById, productId); err != nil {
		return
	}

	if product.Items, err = r.getProductItems(ctx, []int{productId}); err != nil {
		return
	}

	if product.Images, err = r.getImages(ctx, []int{productId}); err != nil {
		return
	}

	if product.Materials, err = r.getMaterials(ctx, []int{productId}); err != nil {
		return
	}

	return
}

func (r *StorageRepositoryImpl) GetProducts(ctx storagecontext.StorageContext, limit int, offset int) (products []Product, err error) {
	err = r.db.SelectContext(ctx.Ctx(), &products, sql.GetProducts, limit, offset)
	if err != nil {
		return nil, err
	}

	for i := range products {
		productId := products[i].Id

		products[i].Items, err = r.getProductItems(ctx, []int{productId})
		if err != nil {
			return nil, err
		}

		products[i].Images, err = r.getImages(ctx, []int{productId})
		if err != nil {
			return nil, err
		}

		products[i].Materials, err = r.getMaterials(ctx, []int{productId})
		if err != nil {
			return nil, err
		}
	}

	return
}

func (r *StorageRepositoryImpl) AddProduct(ctx storagecontext.StorageContext, product Product) error {
	productId, err := r.insertWithReturningId(ctx, sql.AddProduct, product)
	if err != nil {
		return err
	}

	for _, item := range product.Items {
		item.ProductId = productId
		if _, err = r.db.NamedExecContext(ctx.Ctx(), sql.AddProductItems, item); err != nil {
			return err
		}
	}

	for _, material := range product.Materials {
		material.ProductId = productId
		if _, err = r.db.NamedExecContext(ctx.Ctx(), sql.AddProductMaterials, material); err != nil {
			return err
		}
	}

	for _, image := range product.Images {
		image.ProductId = productId
		if _, err = r.db.NamedExecContext(ctx.Ctx(), sql.AddProductImages, image); err != nil {
			return err
		}
	}

	return nil
}

// TODO: ОБНОВЛЕНИЯ СКОРЕЕ ВСЕГО НУЖНО БУДЕТ МЕНЯТЬ ДЛЯ ОСНОВНОГО СЕРВИСА

func (r *StorageRepositoryImpl) UpdateProduct(ctx storagecontext.StorageContext, product Product) error {
	if _, err := r.db.NamedExecContext(ctx.Ctx(), sql.UpdateProduct, product); err != nil {
		return err
	}

	for _, item := range product.Items {
		if _, err := r.db.NamedExecContext(ctx.Ctx(), sql.UpdateProductItems, item); err != nil {
			return err
		}
	}

	return nil
}

func (r *StorageRepositoryImpl) DeleteProduct(ctx storagecontext.StorageContext, productId int) error {
	if _, err := r.db.ExecContext(ctx.Ctx(), sql.DeleteProductImages, productId); err != nil {
		return err
	}

	if _, err := r.db.ExecContext(ctx.Ctx(), sql.DeleteProductMaterials, productId); err != nil {
		return err
	}

	if _, err := r.db.ExecContext(ctx.Ctx(), sql.DeleteProductItems, productId); err != nil {
		return err
	}

	if _, err := r.db.ExecContext(ctx.Ctx(), sql.DeleteProduct, productId); err != nil {
		return err
	}

	return nil
}

func (r *StorageRepositoryImpl) GetFactories(ctx storagecontext.StorageContext, factoryIds []int) (factories []Factory, err error) {
	err = r.db.SelectContext(ctx.Ctx(), &factories, sql.GetFactories, pq.Array(factoryIds))
	return
}

func (r *StorageRepositoryImpl) GetBrands(ctx storagecontext.StorageContext, brandIds []int) (brands []Brand, err error) {
	err = r.db.SelectContext(ctx.Ctx(), &brands, sql.GetBrands, pq.Array(brandIds))
	return
}

func (r *StorageRepositoryImpl) InsertForTest() error {
	_, err := r.db.ExecContext(context.Background(), sql.InsertsForTest)
	return err
}

func (r *StorageRepositoryImpl) getImages(ctx storagecontext.StorageContext, productIds []int) (images []ProductImage, err error) {
	err = r.db.SelectContext(ctx.Ctx(), &images, sql.GetProductImages, pq.Array(productIds))
	return
}

func (r *StorageRepositoryImpl) getMaterials(ctx storagecontext.StorageContext, productIds []int) (materials []ProductMaterial, err error) {
	err = r.db.SelectContext(ctx.Ctx(), &materials, sql.GetProductMaterials, pq.Array(productIds))
	return
}

func (r *StorageRepositoryImpl) getProductItems(ctx storagecontext.StorageContext, productIds []int) (items []ProductItem, err error) {
	err = r.db.SelectContext(ctx.Ctx(), &items, sql.GetProductItems, pq.Array(productIds))
	return
}

func (r *StorageRepositoryImpl) insertWithReturningId(ctx storagecontext.StorageContext, script string, item any) (id int, err error) {
	query, args, err := r.db.BindNamed(script, item)
	if err != nil {
		return 0, err
	}

	return id, r.db.GetContext(ctx.Ctx(), &id, query, args...)
}

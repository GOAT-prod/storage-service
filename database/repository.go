package database

import (
	"context"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	"storage-service/database/sql"
	"storage-service/tools/storagecontext"
)

type StorageRepository interface {
	GetAllProducts(ctx storagecontext.StorageContext, limit int, offset int) ([]Product, error)
	GetImages(ctx storagecontext.StorageContext, productIds []int) ([]ProductImage, error)
	GetMaterials(ctx storagecontext.StorageContext, productIds []int) ([]ProductMaterial, error)
	GetProductItems(ctx storagecontext.StorageContext, productIds []int) (items []ProductItem, err error)
	GetFactories(ctx storagecontext.StorageContext, factoryIds []int) (factories []Factory, err error)
	GetBrands(ctx storagecontext.StorageContext, brandIds []int) (brands []Brand, err error)

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

func (r *StorageRepositoryImpl) GetAllProducts(ctx storagecontext.StorageContext, limit int, offset int) (info []Product, err error) {
	err = r.db.SelectContext(ctx.Ctx(), &info, sql.GetProducts, limit, offset)
	return
}

func (r *StorageRepositoryImpl) GetImages(ctx storagecontext.StorageContext, productIds []int) (images []ProductImage, err error) {
	err = r.db.SelectContext(ctx.Ctx(), &images, sql.GetProductImages, pq.Array(productIds))
	return
}

func (r *StorageRepositoryImpl) GetMaterials(ctx storagecontext.StorageContext, productIds []int) (materials []ProductMaterial, err error) {
	err = r.db.SelectContext(ctx.Ctx(), &materials, sql.GetProductMaterials, pq.Array(productIds))
	return
}

func (r *StorageRepositoryImpl) GetProductItems(ctx storagecontext.StorageContext, productIds []int) (items []ProductItem, err error) {
	err = r.db.SelectContext(ctx.Ctx(), &items, sql.GetProductItems, pq.Array(productIds))
	return
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

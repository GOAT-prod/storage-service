package database

import (
	"context"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	"storage-service/database/sql"
	"storage-service/tools/storagecontext"
)

type StorageRepository interface {
	GetAllProducts(ctx storagecontext.StorageContext) ([]DbProductInfo, error)
	GetLimitProducts(ctx storagecontext.StorageContext, limit int) ([]DbProductInfo, error)
	GetImages(ctx storagecontext.StorageContext, productIds []int) ([]DbImage, error)
	GetMaterials(ctx storagecontext.StorageContext, productIds []int) ([]DbMaterial, error)

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

func (r *StorageRepositoryImpl) GetAllProducts(ctx storagecontext.StorageContext) (info []DbProductInfo, err error) {
	err = r.db.SelectContext(ctx.Ctx(), &info, sql.GetALlProductsInfo)
	return
}

func (r *StorageRepositoryImpl) GetLimitProducts(ctx storagecontext.StorageContext, limit int) (info []DbProductInfo, err error) {
	err = r.db.SelectContext(ctx.Ctx(), &info, sql.GetLimitProductsInfo, limit)
	return
}

func (r *StorageRepositoryImpl) GetImages(ctx storagecontext.StorageContext, productIds []int) (images []DbImage, err error) {
	err = r.db.SelectContext(ctx.Ctx(), &images, sql.GetProductImages, pq.Array(productIds))
	return
}

func (r *StorageRepositoryImpl) GetMaterials(ctx storagecontext.StorageContext, productIds []int) (materials []DbMaterial, err error) {
	err = r.db.SelectContext(ctx.Ctx(), &materials, sql.GetProductMaterials, pq.Array(productIds))
	return
}

func (r *StorageRepositoryImpl) InsertForTest() error {
	_, err := r.db.ExecContext(context.Background(), sql.InsertsForTest)
	return err
}

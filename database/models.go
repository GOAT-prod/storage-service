package database

import "github.com/shopspring/decimal"

type Product struct {
	Id          int             `db:"id"`
	BrandId     int             `db:"brand_id"`
	FactoryId   int             `db:"factory_id"`
	Name        string          `db:"name"`
	Description string          `db:"description"`
	Price       decimal.Decimal `db:"price"`
	Items       []ProductItem
	Materials   []ProductMaterial
	Images      []ProductImage
}

type ProductItem struct {
	Id         int             `db:"id"`
	ProductId  int             `db:"product_id"`
	StockCount int             `db:"stock_count"`
	Size       int             `db:"size"`
	Weight     decimal.Decimal `db:"weight"`
	Color      string          `db:"color"`
}

type ProductMaterial struct {
	Id        int    `db:"id"`
	ProductId int    `db:"product_id"`
	Name      string `db:"name"`
}

type ProductImage struct {
	Id        int    `db:"id"`
	ProductId int    `db:"product_id"`
	ImageUrl  string `db:"url"`
}

type Brand struct {
	Id   int    `db:"id"`
	Name string `db:"name"`
}

type Factory struct {
	Id      int    `db:"id"`
	Name    string `db:"name"`
	Country string `db:"country"`
	City    string `db:"city"`
	Address string `db:"address"`
}

package database

import "github.com/shopspring/decimal"

type DbProductInfo struct {
	Id           int             `db:"Id"`
	Name         string          `db:"Name"`
	Price        decimal.Decimal `db:"Price"`
	Discount     decimal.Decimal `db:"Discount"`
	Size         int             `db:"Size"`
	Color        string          `db:"Color"`
	TypeId       int             `db:"TypeId"`
	TypeName     string          `db:"TypeName"`
	CategoryId   int             `db:"CategoryId"`
	CategoryName string          `db:"CategoryName"`
}

type DbImage struct {
	ProductId int    `db:"product_id"`
	Url       string `db:"url"`
}

type DbMaterial struct {
	ProductId int    `db:"product_id"`
	Name      string `db:"name"`
}

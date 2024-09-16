package domain

import "github.com/shopspring/decimal"

type ProductInfo struct {
	Id        int             `json:"Id"`
	Name      string          `json:"Name"`
	Price     decimal.Decimal `json:"Price"`
	Discount  decimal.Decimal `json:"Discount"`
	Size      int             `json:"Size"`
	Color     string          `json:"Color"`
	Type      ClothesType     `json:"Type"`
	Category  ClothesCategory `json:"Category"`
	Images    []string        `json:"Images"`
	Materials []string        `json:"Materials"`
}

type ClothesType struct {
	Id   int    `json:"Id"`
	Name string `json:"Name"`
}

type ClothesCategory struct {
	Id   int    `json:"Id"`
	Name string `json:"Name"`
}

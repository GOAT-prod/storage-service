package domain

import "github.com/shopspring/decimal"

type Product struct {
	Id          int               `json:"id"`          // Id продукта
	Brand       Brand             `json:"brand"`       // Бренд кроссовок
	Factory     Factory           `json:"factory"`     // Завод изготовитель
	Name        string            `json:"name"`        // Название модели кроссовка
	Description string            `json:"description"` // Описание модели кроссовка
	Price       decimal.Decimal   `json:"price"`       // Цена продукта
	Items       []ProductItem     `json:"items"`       // Варианты кроссовок
	Materials   []ProductMaterial `json:"materials"`   // Материалы изготовления
	Images      []ProductImage    `json:"images"`      // Картинки
}

type ProductItem struct {
	Id         int             `json:"id"`         // Id варианта кроссовка
	StockCount int             `json:"stockCount"` // Кол-во на складе
	Size       int             `json:"size"`       // Размер
	Weight     decimal.Decimal `json:"weight"`     // Вес
	Color      string          `json:"color"`      // Цвет
}

type Brand struct {
	Id   int    `json:"id"`   // Id бренда кроссовок
	Name string `json:"name"` // Наименование бренда
}

type Factory struct {
	Id      int    `json:"id"`      // Id завода изготовителя
	Name    string `json:"name"`    // Название завода
	Country string `json:"country"` // Страна
	City    string `json:"city"`    // Город
	Address string `json:"address"` // Конкретный адрес
}

type ProductImage struct {
	Id       int    `json:"id"`       // Id картинки
	ImageUrl string `json:"imageUrl"` // Ссылка на ресурс с картинкой
}

type ProductMaterial struct {
	Id   int    `json:"id"`   // Id материала
	Name string `json:"name"` // Название материала
}

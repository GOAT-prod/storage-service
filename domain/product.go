package domain

import "github.com/shopspring/decimal"

type Product struct {
	Id          int               `json:"Id"`          // Id продукта
	Brand       Brand             `json:"Brand"`       // Бренд кроссовок
	Factory     Factory           `json:"Factory"`     // Завод изготовитель
	Name        string            `json:"Name"`        // Название модели кроссовка
	Description string            `json:"Description"` // Описание модели кроссовка
	Price       decimal.Decimal   `json:"Price"`       // Цена продукта
	Items       []ProductItem     `json:"Items"`       // Варианты кроссовок
	Materials   []ProductMaterial `json:"Materials"`   // Материалы изготовления
	Images      []ProductImage    `json:"Images"`      // Картинки
}

type ProductItem struct {
	Id         int             `json:"Id"`         // Id варианта кроссовка
	StockCount int             `json:"StockCount"` // Кол-во на складе
	Size       int             `json:"Size"`       // Размер
	Weight     decimal.Decimal `json:"Weight"`     // Вес
	Color      string          `json:"Color"`      // Цвет
}

type Brand struct {
	Id   int    `json:"Id"`   // Id бренда кроссовок
	Name string `json:"Name"` // Наименование бренда
}

type Factory struct {
	Id      int    `json:"Id"`      // Id завода изготовителя
	Name    string `json:"Name"`    // Название завода
	Country string `json:"Country"` // Страна
	City    string `json:"City"`    // Город
	Address string `json:"address"` // Конкретный адрес
}

type ProductImage struct {
	Id       int    `json:"Id"`       // Id картинки
	ImageUrl string `json:"ImageUrl"` // Ссылка на ресурс с картинкой
}

type ProductMaterial struct {
	Id   int    `json:"Id"`   // Id материала
	Name string `json:"Name"` // Название материала
}

package sql

import _ "embed"

var (
	//go:embed get_products.sql
	GetProducts string

	//go:embed get_images.sql
	GetProductImages string

	//go:embed get_materials.sql
	GetProductMaterials string

	//go:embed get_product_items.sql
	GetProductItems string

	//go:embed get_factories.sql
	GetFactories string

	//go:embed get_brands.sql
	GetBrands string

	//go:embed mocks/for_local_test.sql
	InsertsForTest string
)

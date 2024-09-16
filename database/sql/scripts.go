package sql

import _ "embed"

var (
	//go:embed get_product_info.sql
	GetALlProductsInfo string

	//go:embed get_limit_product_info.sql
	GetLimitProductsInfo string

	//go:embed get_product_images.sql
	GetProductImages string

	//go:embed get_product_materials.sql
	GetProductMaterials string

	//go:embed mocks/for_local_test.sql
	InsertsForTest string
)

package sql

import _ "embed"

var (
	//go:embed get_product_by_id.sql
	GetProductById string

	//go:embed get_products.sql
	GetProducts string

	//go:embed add_product.sql
	AddProduct string

	//go:embed update_product.sql
	UpdateProduct string

	//go:embed delete_product.sql
	DeleteProduct string

	//go:embed get_images.sql
	GetProductImages string

	//go:embed add_product_images.sql
	AddProductImages string

	//go:embed delete_images.sql
	DeleteProductImages string

	//go:embed get_materials.sql
	GetProductMaterials string

	//go:embed add_product_material.sql
	AddProductMaterials string

	//go:embed delete_product_materials.sql
	DeleteProductMaterials string

	//go:embed get_product_items.sql
	GetProductItems string

	//go:embed add_product_item.sql
	AddProductItems string

	//go:embed update_product_item.sql
	UpdateProductItems string

	//go:embed delete_product_item.sql
	DeleteProductItems string

	//go:embed get_factories.sql
	GetFactories string

	//go:embed add_factory.sql
	AddFactories string

	//go:embed get_brands.sql
	GetBrands string

	//go:embed add_brand.sql
	AddBrands string

	//go:embed mocks/for_local_test.sql
	InsertsForTest string
)

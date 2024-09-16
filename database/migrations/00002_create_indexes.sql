-- +goose Up
create index if not exists products_type_id_idx on products(type_id);
create index if not exists products_category_id_idx on products(category_id);
create index if not exists products_material_id_ids on products(material_id);
create index if not exists images_product_id_idx on images(product_id);

-- +goose Down
drop index if exists products_type_id_idx;
drop index if exists products_category_id_idx;
drop index if exists products_material_id_ids;
drop index if exists images_product_id_idx;
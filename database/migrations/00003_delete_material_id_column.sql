-- +goose Up
drop index if exists products_material_id_ids;

alter table products
drop column if exists material_id;

create table if not exists product_materials
(
    id          serial primary key,
    product_id  int references products (id),
    material_id int references materials (id)
);

create index if not exists product_materials_product_id_ids on product_materials (product_id);
create index if not exists product_materials_material_id_idx on product_materials (material_id);

-- +goose Down
drop index if exists product_materials_material_id_idx;
drop index if exists product_materials_product_id_ids;
drop table if exists product_materials;
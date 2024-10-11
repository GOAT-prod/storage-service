insert into product (brand_id, factory_id, name, description, price, is_approved, is_deleted)
values (:brand_id, :factory_id, :name, :description, :price, :is_approved, :is_deleted)
returning id
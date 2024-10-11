update product
set brand_id = :brand_id,
    factory_id = :factory_id,
    name = :name,
    description = :description,
    price = :price,
    is_approved = :is_approved,
    is_deleted = :is_deleted
where id = :id;

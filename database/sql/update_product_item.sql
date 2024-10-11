update product_item
set product_id = :product_id,
    stock_count = :stock_count,
    size = :size,
    weight = :weight,
    color = :color
where id = :id;

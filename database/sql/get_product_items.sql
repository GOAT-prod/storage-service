select id,
       product_id,
       stock_count,
       size,
       weight,
       color
from product_item
where product_id = any ($1)
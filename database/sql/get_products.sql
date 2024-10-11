select id,
       brand_id,
       factory_id,
       name,
       description,
       price
from product
limit $1
offset $2;

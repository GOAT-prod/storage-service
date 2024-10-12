select id,
       brand_id,
       factory_id,
       name,
       description,
       price
from product
where id = $1
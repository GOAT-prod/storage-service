select url,
       product_id
from images
where product_id = any($1)
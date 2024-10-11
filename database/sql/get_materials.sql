select m.id,
       pm.product_id,
       m.name
from product_material pm
         join material m on pm.material_id = m.id
where pm.product_id = any ($1)
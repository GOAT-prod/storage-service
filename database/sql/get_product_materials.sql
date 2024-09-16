select pm.product_id,
       m.name
from product_materials pm
         join materials m on pm.material_id = m.id
where pm.product_id = any ($1)
select p.id       as "Id",
       p.name     as "Name",
       p.price    as "Price",
       p.discount as "Discount",
       p.size     as "Size",
       p.color    as "Color",
       ct.id      as "TypeId",
       ct.name    as "TypeName",
       cc.id      as "CategoryId",
       cc.name    as "CategoryName"
from products p
         join clothes_type ct on p.type_id = ct.id
         join clothes_categories cc on p.category_id = cc.id
limit $1
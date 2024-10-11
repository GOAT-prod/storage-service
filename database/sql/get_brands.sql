select id,
       name
from brand
where id = any ($1)
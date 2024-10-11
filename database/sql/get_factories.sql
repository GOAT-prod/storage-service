select id,
       name,
       country,
       city,
       address
from factory
where id = any ($1)
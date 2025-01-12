SHOW DATABASES;
USE uas_ppb;

SHOW TABLES;
select *
from users;


select *
from products;


select *
from users;
select *
from carts;


SELECT p.name,
       ci.jumlah,
       p.price,
       ci.subtotal,
       ci.deleted_at as buy_time
FROM cart_items ci
         LEFT JOIN products p ON ci.product_id = p.id
WHERE ci.deleted_at IS NOT NULL;

select *
from cart_items;

select *
from products;

select * from users;
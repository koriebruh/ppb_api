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

select *
from users;


show tables;
select *
from cart_items;
select *
from carts;
select *
from products;
select *
from shippings;
select *
from users;

SELECT *
FROM carts c
         right JOIN cart_items ci on c.id = ci.cart_id
where c.user_id = 2;


SELECT p.name,
       ci.jumlah,
       p.price,
       ci.subtotal,
       ci.deleted_at as buy_time
FROM carts c
         RIGHT JOIN cart_items ci ON c.id = ci.cart_id
         LEFT JOIN products p ON ci.product_id = p.id
where c.user_id = 1 and ci.deleted_at IS NOT NULL;


SELECT p.name,
       ci.jumlah,
       p.price,
       ci.subtotal,
       ci.deleted_at as buy_time
FROM carts c
         RIGHT JOIN cart_items ci ON c.id = ci.cart_id
         LEFT JOIN products p ON ci.product_id = p.id
WHERE ci.deleted_at BETWEEN '2025-01-01' AND '2025-01-02';

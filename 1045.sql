select
    customer_id
from 
    Customer as c
group by
    customer_id
having
    count(distinct product_key) = (
        select count(product_key)
        from Product
    );
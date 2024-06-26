select
    p.product_name,
    sum(o.unit) as unit
from
    Products as p
inner join
    Orders as o
        on p.product_id = o.product_id and
        extract(year from order_date) = '2020' and
        extract(month from order_date) = '02'
group by
    p.product_id, p.product_name
having
    sum(o.unit) >= 100
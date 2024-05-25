-- Write your PostgreSQL query statement below
select 
    p.product_id,
    coalesce(round(sum(u.units * p.price) / sum(u.units)::numeric, 2), 0) as average_price
from 
    Prices as p
left join
    UnitsSold as u
on
    p.start_date <= u.purchase_date 
    and u.purchase_date <= p.end_date
    and p.product_id = u.product_id
group by
    p.product_id
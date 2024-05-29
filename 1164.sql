-- Write your PostgreSQL query statement below
with Dates_Til_bounder as (
    select
        *
    from 
        Products
    where
        change_date <= '2019-08-16'::date
),

maxDate as (
    select
        product_id,
        max(change_date) as change_date
    from 
        Dates_Til_bounder
    group by
        product_id
)

select
    distinct p.product_id,
    case
        when m.product_id is null then 10
    else
        p.new_price
    end as price
from
    Products as p
left join
    maxDate as m
        on m.product_id = p.product_id and m.change_date = p.change_date
where
    p.product_id not in (select product_id from maxDate) 
    or m.product_id is not null
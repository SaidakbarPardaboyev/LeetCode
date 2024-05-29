-- Write your PostgreSQL query statement below
with first_date_sold_products as (
    select
        product_id,
        min(year) as year
    from
        Sales
    group by
        product_id
)

select
    f.product_id,
    f.year as first_year,
    quantity,
    price
from
    Sales as s
inner join
    first_date_sold_products as f
        on f.product_id = s.product_id and f.year = s.year
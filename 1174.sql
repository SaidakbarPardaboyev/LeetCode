-- Write your PostgreSQL query statement below
with firstOrder as (
    select
        customer_id,
        min(order_date) as first_order_date
    from 
        Delivery
    group by
        customer_id
)

select
    round(count(
        case
            when d.order_date = d.customer_pref_delivery_date then 1
        end
    )::numeric / count(distinct d.customer_id) * 100, 2) as immediate_percentage
from
    Delivery as d
inner join
    firstOrder as f
        on f.customer_id = d.customer_id
        and d.order_date = f.first_order_date
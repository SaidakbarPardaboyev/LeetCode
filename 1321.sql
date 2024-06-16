-- Write your PostgreSQL query statement below
with minDate as (
    select
        min(visited_on) + 6 as minn
    from
        customer
),

distinctDate as (
    select
        distinct visited_on
    from
        customer
)

select
    c1.visited_on,
    sum(c2.amount) as amount,
    round(sum(c2.amount) / count(distinct c2.visited_on)::numeric, 2) as average_amount
from
    distinctDate as c1
inner join
    customer as c2
        on c2.visited_on <= c1.visited_on and 
        c1.visited_on - c2.visited_on <= 6
group by
    c1.visited_on
having
    c1.visited_on >= (select minn from mindate)
-- Write your PostgreSQL query statement below
with counts as (select
        requester_id as id
    from
        requestaccepted
    union all
    select
        accepter_id as id
    from
        requestaccepted
)

select
    id,
    count(id) as num
from 
    counts
group by
    id
order by
    count(id) desc
limit 1
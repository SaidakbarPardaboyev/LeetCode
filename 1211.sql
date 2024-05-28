-- # Write your MySQL query statement below
select
    query_name,
    round(avg(rating::numeric / position), 2) as quality,
    round(sum(case when rating < 3 then 1 else 0 end)::numeric / count(*) * 100, 2) as poor_query_percentage
from
    Queries
group by
    query_name
having
    query_name is not null
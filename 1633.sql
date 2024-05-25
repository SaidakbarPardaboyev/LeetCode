-- Write your PostgreSQL query statement below
select 
    r.contest_id,
    round(count(r.user_id) / (select count(*) from Users)::numeric * 100, 2) as percentage
from Register as r
group by r.contest_id
order by percentage desc, r.contest_id asc
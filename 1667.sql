-- Write your PostgreSQL query statement below
select
    user_id,
    upper(substring(name from 1 for 1)) || lower(substring(name from 2 for char_length(name)-1)) as name
from
    Users
order by
    user_id
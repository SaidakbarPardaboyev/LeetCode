-- Write your PostgreSQL query statement below
select
    activity_date as day,
    count(distinct user_id) as active_users
from
    Activity as a
group by
    activity_date
having
    activity_date > '2019-07-27'::date - 30 and activity_date <= '2019-07-27'::date
-- Write your PostgreSQL query statement below
select
    distinct l.num as ConsecutiveNums
from
    Logs as l
left join
    Logs as l2
        on l.id - l2.id = -1 and l.num = l2.num
left join
    Logs as l3
        on l2.id - l3.id = -1 and l3.num = l2.num
where
    l2.id is not null and l3.id is not null
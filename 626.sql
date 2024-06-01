-- Write your PostgreSQL query statement below
select
    s1.id,
    case
        when s3.id is not null then s3.student
        when s2.id is not null then s2.student
        else s1.student
    end as student
from 
    Seat as s1
left join
    Seat as s2
        on s1.id = s2.id-1 and s2.id % 2 = 0
left join
    Seat as s3
        on s1.id = s3.id+1 and s3.id % 2 != 0
order by
    s1.id
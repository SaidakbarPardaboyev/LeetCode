-- Write your PostgreSQL query statement below
select
    q.person_name
from 
    Queue as q
inner join
    Queue as q1
        on q1.turn <= q.turn
group by
    q.turn, q.person_name
having
    sum(q1.weight) <= 1000
order by
    q.turn desc
limit 1

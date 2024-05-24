-- Write your PostgreSQL query statement below
select a1.machine_id, round((AVG(a1.timestamp - a2.timestamp))::numeric,3) as processing_time
from Activity as a1
inner join Activity as a2
on a1.machine_id = a2.machine_id and a1.process_id = a2.process_id
and a1.activity_type = 'end' and a2.activity_type = 'start'
group by a1.machine_id
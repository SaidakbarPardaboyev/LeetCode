# Write your MySQL query statement below
select p.project_id, round(sum(e.experience_years) / count(e.employee_id),2) as average_years
from project as p
inner join employee as e
on p.employee_id = e.employee_id and e.experience_years is not null
group by p.project_id
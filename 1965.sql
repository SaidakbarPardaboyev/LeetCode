-- Write your PostgreSQL query statement below
select 
    case 
        when e.employee_id is not null then e.employee_id
        when s.employee_id is not null then s.employee_id
    end as employee_id
from Employees as e
full join Salaries as s
on e.employee_id = s.employee_id
where e.employee_id is null or s.employee_id is null
order by employee_id
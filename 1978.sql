-- Write your PostgreSQL query statement below
select e.employee_id
from Employees as e
where e.manager_id is not null
and e.manager_id not in (
    select employee_id
    from Employees
) and e.salary < 30000
order by e.employee_id
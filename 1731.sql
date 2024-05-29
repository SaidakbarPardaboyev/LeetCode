-- Write your PostgreSQL query statement below
select
    e.employee_id,
    e.name,
    count(*) as reports_count,
    avg(r.age)::int as average_age
from
    Employees as e
inner join
    Employees as r
        on e.employee_id = r.reports_to
group by
    e.employee_id, e.name
order by
    e.employee_id
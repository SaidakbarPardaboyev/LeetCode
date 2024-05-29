-- Write your PostgreSQL query statement below
with prime as (
    select
        employee_id,
        department_id
    from
        Employee
    where
        primary_flag = 'Y'
),

counts as (
    select
        employee_id,
        count(*) as counts
    from 
        Employee
    group by
        employee_id
)

select
    distinct e.employee_id,
    case
        when p.employee_id is null then e.department_id
    else 
        p.department_id 
    end as department_id
from 
    Employee as e
left join
    prime as p
        on p.employee_id = e.employee_id
right join
    counts as c
        on c.employee_id = e.employee_id
where c.counts <= 1 or p.department_id is not null
-- -- Write your PostgreSQL query statement below
with uniqueEmployeeinfo as (
    select
        id,
        name,
        salary,
        departmentId
    from 
        employee as e
    where 
        e.salary not in (
            select
                salary
            from
                employee as e1
            where
                e1.id < e.id and 
                e1.departmentId = e.departmentId
        )
)

select
    d.name as Department,
    e.name as Employee,
    salary
from
    employee as e
inner join 
    department as d
        on d.id = e.departmentId
where
    e.salary >= COALESCE(
        (select
                    em.Salary
                from 
                    uniqueEmployeeinfo as em
                inner join
                    department as dp
                        on dp.id = em.departmentId and
                            dp.id = d.id
                order by
                    dp.id asc,
                    em.Salary desc
                offset 
                    2
                limit
                    1
        ), 0)
order by
    d.id
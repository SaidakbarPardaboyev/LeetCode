-- Write your PostgreSQL query statement below
select e.unique_id, em.name from Employees as em
left join EmployeeUNI as e
on e.id = em.id; 
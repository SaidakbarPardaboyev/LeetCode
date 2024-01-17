# Write your MySQL query statement below
select s.name
from orders as o
left join company as c
on o.com_id = c.com_id
right join salesPerson as s
on s.sales_id = o.sales_id and c.name = "RED"
where o.order_id is null
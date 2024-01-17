# Write your MySQL query statement below
select c.name as customers from Customers as c left join Orders as o on o.customerId = c.id where o.customerId is null;
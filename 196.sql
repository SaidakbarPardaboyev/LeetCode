# Write your MySQL query statement below
delete p2 from Person as p1 inner join Person as p2 on p1.email = p2.email where p1.id < p2.id;
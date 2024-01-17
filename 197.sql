# Write your MySQL query statement below
select w2.id as Id
from Weather as w1
join Weather as w2 on w1.temperature < w2.temperature and w1.recordDate + interval 1 day = w2.recordDate;
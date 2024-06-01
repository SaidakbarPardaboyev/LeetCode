-- Write your PostgreSQL query statement below
select
    *
from
    Users as s
where
    s.mail like '%@leetcode.com' and 
    left(mail, length(mail)-13) ~ '^[a-zA-Z0-9_.-]+$' and 
    substring(s.mail from 1 for 1) ~ '^[a-zA-Z]'
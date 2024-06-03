-- Write your PostgreSQL query statement below
(select
    name as results
from
    movies as m
inner join
    MovieRating as mr
        on m.movie_id = mr.movie_id
inner join
    users as u
        on u.user_id = mr.user_id
group by
    name
order by
    count(name) desc, name
limit 1)

union all

(select
    title as results
from
    movies as m
inner join
    MovieRating as mr
        on m.movie_id = mr.movie_id and 
        extract(year from mr.created_at) = 2020 and
        extract(month from mr.created_at) = 2
inner join
    users as u
        on u.user_id = mr.user_id
group by
    m.movie_id, title
order by
    avg(rating) desc, title
limit 1)
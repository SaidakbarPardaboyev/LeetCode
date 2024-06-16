-- Write your PostgreSQL query statement below
with tiv2015 as (select
    distinct i1.pid,
    i1.tiv_2015,
    i1.tiv_2016,
    i1.lat,
    i1.lon
from
    insurance as i1
inner join 
    insurance as i2
        on i1.tiv_2015 = i2.tiv_2015 and
            i1.pid != i2.pid
),

uniqueLokations as (select
    i1.pid,
    i1.tiv_2015,
    i1.tiv_2016,
    i1.lat,
    i1.lon
from
    insurance as i1
left join
    insurance as i2
        on i1.lat = i2.lat and i1.lon = i2.lon and
            i1.pid != i2.pid
where
    i2.lat is null
)

select 
    round(sum(t.tiv_2016)::numeric, 2) as tiv_2016
from
    tiv2015 as t
inner join
    uniqueLokations as l
        on t.pid = l.pid

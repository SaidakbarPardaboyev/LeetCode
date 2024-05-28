-- Write your PostgreSQL query statement below
with first_log_date_of_user as (
    select
        player_id,
        min(event_date) as event_date
    from
        Activity
    group by
        player_id
)

select 
    round(count(f.player_id)::numeric / count(distinct a.player_id), 2) as fraction
from
    Activity as a
left join
    first_log_date_of_user as f
        on f.player_id = a.player_id 
        and a.event_date-1 = f.event_date
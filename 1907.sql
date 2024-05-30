select
    unnest(array['Low Salary', 'Average Salary', 'High Salary']) as category,
    unnest(array[
        sum(case when income < 20000 then 1 else 0 end),
        sum(case when income >= 20000 and income <= 50000 then 1 else 0 end),
        sum(case when income > 50000 then 1 else 0 end)
            
    ]) as accounts_count
from Accounts
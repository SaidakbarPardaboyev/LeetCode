class Solution:
    def average(self, salary: list[int]) -> float:
        salary.remove(max(salary))
        salary.remove(min(salary))
        return sum(salary) / len(salary)
    
print(Solution().average([4000,3000,1000,2000]))
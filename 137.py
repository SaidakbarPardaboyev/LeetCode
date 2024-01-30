class Solution:
    def singleNumber(self, nums: list[int]) -> int:
        for i in set(nums):
            if nums.count(i) == 1:
                return i
            

print(Solution().singleNumber([1,0,1,0,1,0,99]))
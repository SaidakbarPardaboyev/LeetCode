class Solution:
    def singleNumber(self, nums: list[int]) -> list[int]:
        res = list(set(nums))
        res.sort(key=lambda x: nums.count(x))
        return [res[0], res[1]]
    
print(Solution().singleNumber([1,2,1,3,2,5]))
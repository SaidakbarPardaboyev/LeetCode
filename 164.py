class Solution:
    def maximumGap(self, nums: list[int]) -> int:
        nums.sort()
        if not len(nums) >= 2:
            return 0

        res = 0
        for i in range(1, len(nums)):
            tem = nums[i] - nums[i-1]
            res = max(res, tem)
        
        return res
    
print(Solution().maximumGap([3,6,9,1]))
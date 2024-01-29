class Solution:
    def longestConsecutive(self, nums: list[int]) -> int:
        nums = list(set(nums))
        nums.sort()
        
        if len(nums) == 0:
            return 0

        res = 0
        tem = 0

        for i in range(len(nums)-1):
            if nums[i] + 1 == nums[i+1]:
                tem += 1
            else:
                res = max(tem, res)
                tem = 0
            
        res = max(tem, res)
        res += 1
        return res
    
print(Solution().longestConsecutive([100,4,200,1,3,2]))
class Solution:
    def missingNumber(self, nums: list[int]) -> int:
        num = 0
        for i in range(len(nums)):
            if num not in nums:
                return num
            else:
                num += 1
        return len(nums)
    
print(Solution().missingNumber([3,0,1]))
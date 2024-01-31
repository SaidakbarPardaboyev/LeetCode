class Solution:
    def findPeakElement(self, nums: list[int]) -> int:
        for i in range(1, len(nums)-1):
            if nums[i] > nums[i-1] and nums[i] > nums[i+1]:
                return i
            
print(Solution().findPeakElement([1,2,3,1]))
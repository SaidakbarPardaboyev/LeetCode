class Solution:
    def findMin(self, nums: list[int]) -> int:
        for i in range(len(nums)):
            box = nums[i]
            j = i-1
            while j >= 0 and nums[j] > box:
                if nums[j] > box:
                    nums[j+1] = nums[j]
                j -= 1
            nums[j+1] = box
        
        return nums[0]
                
print(Solution().findMin([3,4,5,1,2]))
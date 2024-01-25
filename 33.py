class Solution:
    def search(self, nums: list[int], target: int) -> int:
        res = -1

        for i in range(len(nums)):
            if nums[i] == target:
                return i
                
        return res
    
print(Solution().search([4,5,6,7,0,1,2], target = 0))
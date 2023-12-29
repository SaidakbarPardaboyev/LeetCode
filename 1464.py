class Solution:
    def maxProduct(self, nums: list[int]) -> int:
        max1 = max(nums)
        nums.remove(max1)
        max2 = max(nums)

        return (max1 - 1) * (max2 - 1)
    
result = Solution()
print(result.maxProduct([3,4,5,2]))
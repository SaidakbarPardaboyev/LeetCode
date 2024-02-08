class Solution:
    def rotate(self, nums: list[int], k: int) -> None:
        """
        Do not return anything, modify nums in-place instead.
        """
        i = k % len(nums)
        if i == 0:
            return nums
        
        nums[:] = nums[-i:] + nums[:-i]
        return nums

print(Solution().rotate([1,2,3,4,5,6,7], k = 3))
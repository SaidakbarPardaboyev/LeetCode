class Solution:
    def sortColors(self, nums: list[int]) -> None:
        for i in range(1, len(nums)):
            box = nums[i]
            j = i-1
            while j >= 0 and nums[j] > box:
                if nums[j] > box:
                    nums[j+1] = nums[j]
                j -= 1
            nums[j+1] = box
        return nums


print(Solution().sortColors([2,0,2,1,1,0]))
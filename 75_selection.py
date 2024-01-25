class Solution:
    def sortColors(self, nums: list[int]) -> None:
        leng = len(nums)
        for i in range(leng):
            ind = i
            for j in range(i+1, leng):
                if nums[j] < nums[ind]:
                    ind = j
            nums[i], nums[ind] = nums[ind], nums[i]
        return nums


print(Solution().sortColors([2,0,2,1,1,0]))
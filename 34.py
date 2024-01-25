class Solution:
    def searchRange(self, nums: list[int], target: int) -> list[int]:
        ls = [-1,-1]

        try:
            ls[0] = nums.index(target)
        except:
            pass

        for i in range(len(nums)-1, -1, -1):
            if nums[i] == target:
                ls[1] = i
                break

        return ls
    
print(Solution().searchRange([5,7,7,8,8,10], target = 8))
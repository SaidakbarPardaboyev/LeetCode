class Solution:
    def increasingTriplet(self, nums: list[int]) -> bool:
        for i in range(0, len(nums)-2):
            for j in range(i+1, len(nums)-1):
                for k in range(j+1, len(nums)):
                    if nums[i] < nums[j] and nums[j] < nums[k]:
                        return True
        return False

print(Solution().increasingTriplet([1,2,3,4,5]))
class Solution:
    def findPairs(self, nums: list[int], k: int) -> int:
        res = set()
        for i in range(len(nums)):
            for j in range(i+1, len(nums)):
                if abs(nums[i] - nums[j]) == k:
                    if nums[i] > nums[j]:
                        res.add((nums[i], nums[j]))
                    else:
                        res.add((nums[j], nums[i]))
        return len(res)
    
print(Solution().findPairs([1,2,4,4,3,3,0,9,2,3], 3))
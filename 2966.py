class Solution:
    def divideArray(self, nums: list[int], k: int) -> list[list[int]]:
        nums.sort()

        res = list()

        i = 0
        while i < len(nums):
            tem = [nums[i], nums[i+1], nums[i+2]]
            if tem[-1] - tem[0] <= k:
                res.append(tem)
            else:
                return []
            i += 3
        
        return res
    
print(Solution().divideArray([1,3,4,8,7,9,3,5,1], 2))
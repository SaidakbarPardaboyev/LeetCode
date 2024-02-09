class Solution:
    def findDuplicate(self, nums: list[int]) -> int:
        son = sorted(nums)

        for i in range(1, len(son)):
            if son[i] == son[i - 1]:
                return son[i]
            
print(Solution().findDuplicate([1,3,4,2,2]))
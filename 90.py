class Solution:
    def subsetsWithDup(self, nums: list[int]) -> list[list[int]]:
        from itertools import combinations
        nums.sort()
        com = list()
        res = [[]]

        for i in range(len(nums)):
            tem = list(combinations(nums, i+1))
            tem = set(tem)
            com.append(tem)

        for i in com:
            for j in i:
                res.append(list(j))
        return res
        
print(Solution().subsetsWithDup([1,2,2]))
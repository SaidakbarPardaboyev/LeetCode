class Solution:
    def subsets(self, nums: list[int]) -> list[list[int]]:
        res = [[]]
        com = []

        from itertools import combinations
        for i in range(len(nums)):
            tem = combinations(nums, i+1)
            com.append(tem)

        for i in com:
            for j in i:
                res.append(list(j))

        return res


print(Solution().subsets([1,2,3]))
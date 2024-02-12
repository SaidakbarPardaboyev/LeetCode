class Solution:
    def kthSmallest(self, matrix: list[list[int]], k: int) -> int:
        res = list()
        for i in matrix:
            for j in i:
                res.append(j)
        res.sort()
        return res[k-1]
    
print(Solution().kthSmallest([[1,5,9],[10,11,13],[12,13,15]], 8))
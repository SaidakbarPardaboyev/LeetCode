class Solution:
    def combine(self, n, k):
        ls = [i for i in range(1, n+1)]
        from itertools import combinations
        res = list(combinations(ls, k))

        return [list(i) for i in res]

print(Solution().combine(4, 2))
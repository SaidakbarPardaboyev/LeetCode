class Solution:
    def superPow(self, a: int, b: list[int]) -> int:
        res = 1
        for i in b:
            res = pow(res, 10) * pow(a, i) % 1337

        return res
    
print(Solution().superPow(2, [1,0]))
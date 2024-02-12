class Solution:
    def lexicalOrder(self, n: int) -> list[int]:
        res = [i for i in range(1, n+1)]
        res.sort(key=lambda x: str(x))

        return res
    
print(Solution().lexicalOrder(13))
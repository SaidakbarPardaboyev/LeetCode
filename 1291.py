class Solution:
    def sequentialDigits(self, low: int, high: int) -> list[int]:
        res = list()

        st = "123456789"

        for i in range(len(st)):
            for j in range(i+1,len(st)):
                if low <= int(st[i:j+1]) <= high:
                    res.append(int(st[i:j+1]))
        return sorted(res)

print(Solution().sequentialDigits(10, 10000000))

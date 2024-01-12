class Solution:
    def maxLengthBetweenEqualCharacters(self, s: str) -> int:
        tem = s[::-1]
        res = 0
        lamp = 0
        size = len(s)

        for i in s:
            if s.count(i) > 1:
                lamp = 1
                ind = s.find(i)
                ind2 = tem.find(i)
                print(ind, ind2)
                if (size - ind2) - ind - 2 > res:
                    res = (size - ind2) - ind - 2

        if lamp == 0:
            res = -1
        return res

print(Solution().maxLengthBetweenEqualCharacters("dva"))
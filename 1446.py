class Solution:
    def maxPower(self, s: str) -> int:
        res = 0
        temLetter = s[0]
        sch = 0

        for i in s:
            if i == temLetter:
                sch += 1
            else:
                if sch >= res:
                   res = sch
                temLetter = i
                sch = 1
        if sch >= res:
            res = sch
        return res 

result = Solution()
print(result.maxPower("hooraaaaaaaaaaay"))
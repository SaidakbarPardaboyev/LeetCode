class Solution:
    def shiftingLetters(self, s: str, shifts: list[int]) -> str:
        dev = 0
        summa = sum(shifts)
        for i in range(len(shifts)):
            tem = shifts[i]
            shifts[i] = summa - dev
            dev += tem

        for i in range(len(s)):
            tem = (ord(s[i]) + (shifts[i] % 26)) % (ord('z')+1)
            if tem < ord('a'):
                tem += ord('a')
            s = s[:i] + chr(tem) + s[i+1:]

        return s

print(Solution().shiftingLetters("abc", [3,5,9])) # "rpl"
print(Solution().shiftingLetters("bad", [10,20,30])) # "jyh"
print(Solution().shiftingLetters("xrdofd", [70,41,71,72,73,84])) # "surjgj"


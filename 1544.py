class Solution:
    def makeGood(self, s: str) -> str:
        size = len(s)
        bor = 1

        while bor != 0:
            bor = 0
            i = 0
            while i <= size - 2:
                if s[i] != s[i+1]:
                    if s[i].lower() == s[i+1].lower():
                        s = s[:i] + s[i+2:]
                        size -= 2
                        i -= 1
                        bor += 1
                i += 1

        return s
    
print(Solution().makeGood("leEeetcode"))
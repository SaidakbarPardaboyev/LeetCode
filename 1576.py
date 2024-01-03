import random as r

class Solution:
    def modifyString(self, s: str) -> str:
        while s.count('?') != 0:
            ls = ['A']
            ind = s.find('?')
            if ind != 0:
                ls.append(s[ind - 1])
            
            if ind != len(s) - 1:
                ls.append(s[ind + 1])

            n = 'A'

            while n in ls:
                n = chr(r.randint(ord('a'), ord('z')))

            s = s[:ind] + n + s[ind + 1:]    
        return s
    
print(Solution().modifyString("j?qg??b"))
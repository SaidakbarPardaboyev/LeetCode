class Solution:
    def findAnagrams(self, s: str, p: str) -> list[int]:
        res = []
        p = "".join(sorted(p))
        for i in range(len(s)-len(p)+1):
            tem = "".join(sorted(s[i:i+len(p)]))
            if tem == p:
                res.append(i)
        
        return res
    
print(Solution().findAnagrams("cbaebabacd", "abc"))
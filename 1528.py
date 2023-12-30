class Solution:
    def restoreString(self, s: str, indices: list[int]) -> str:
        dic = dict()
        for i in range(len(s)):
            dic.update({indices[i]: s[i]})

        res = list()

        for i in range(len(s)):
            res.append(dic[i])    
                
        return "".join(res)

print(Solution().restoreString("codeleet", [4,5,6,7,0,2,1,3]))
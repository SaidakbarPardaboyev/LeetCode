class Solution:
    def groupAnagrams(self, strs: list[str]) -> list[list[str]]:
        cpy = list(sorted(list(i)) for i in strs)
        res = list()

        i = 0
        while i < len(strs):
            tem = list()
            tem.append(strs[i])
            j = i + 1
            while j < len(strs):
                if cpy[i] == cpy[j]:
                    tem.append(strs[j])
                    cpy.pop(j)
                    strs.pop(j)
                else:
                    j += 1
            cpy.pop(i)
            strs.pop(i)
            res.append(tem)
        return res

print(Solution().groupAnagrams(["eat","tea","tan","ate","nat","bat"]))
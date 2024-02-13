class Solution:
    def frequencySort(self, s: str) -> str:
        dic = dict()

        for i in s:
            if i in dic:
                dic[i] += 1
            else:
                dic[i] = 1

        dic = dict(sorted(dic.items(), key=lambda x: x[1], reverse=True))

        res = ""
        for key, val in dic.items():
            res += (val * key)
    
        return res

print(Solution().frequencySort("tree"))
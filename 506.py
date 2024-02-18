class Solution:
    def findRelativeRanks(self, score: list[int]) -> list[str]:
        ls = score.copy()
        ls.sort(reverse=True)

        dic = dict()
        for i in range(len(ls)):
            if i+1 > 3:
                dic[ls[i]] = str(i+1)
            elif i+1 == 1:
                dic[ls[i]] = "Gold Medal"
            elif i+1 == 2:
                dic[ls[i]] = "Silver Medal"
            elif i+1 == 3:
                dic[ls[i]] = "Bronze Medal"
        
        res = []
        for i in score:
            res.append(dic[i])
        return res

print(Solution().findRelativeRanks([10,3,8,9,4]))
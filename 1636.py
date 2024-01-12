class Solution:
    def frequencySort(self, nums: list[int]) -> list[int]:
        ls = list()
        tem = set(nums)

        for i in tem:
            ls.append([i, nums.count(i)])

        for i in range(len(ls)):
            for j in range(i+1, len(ls)):
                if ls[i][-1] > ls[j][-1]:
                    ls[i], ls[j] = ls[j], ls[i]
                elif ls[i][-1] == ls[j][-1]:
                    if ls[i][0] < ls[j][0]:
                        ls[i], ls[j] = ls[j], ls[i]
        res = list()
        for i in ls:
            while i[-1] != 0:
                res.append(i[0])
                i[-1] -= 1
        return res

print(Solution().frequencySort([1,1,2,2,2,3]))
        
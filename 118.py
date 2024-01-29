class Solution:
    def generate(self, numRows: int) -> list[list[int]]:
        ls = list()
        ls.append([1])

        for i in range(numRows-1):
            tem = list()
            just = ls[-1][-1]
            tem.append(ls[-1][0])
            for j in range(len(ls[-1])-1):
                tem.append(ls[-1][j] + ls[-1][j+1])
            tem.append(just)
            ls.append(tem)

        return ls
    
print(Solution().generate(5))
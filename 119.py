class Solution:
    def getRow(self, rowIndex: int) -> list[int]:
        ls = list()
        ls.append([1])
        tem = list()
        tem.append(1)
        for i in range(rowIndex):
            tem = list()
            just = ls[-1][-1]
            tem.append(ls[-1][0])
            for j in range(len(ls[-1])-1):
                tem.append(ls[-1][j] + ls[-1][j+1])
            tem.append(just)
            ls.append(tem)

        return tem
    
print(Solution().getRow(3))
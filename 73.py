class Solution:
    def setZeroes(self, matrix: list[list[int]]) -> None:
        """
        Do not return anything, modify matrix in-place instead.
        """
        ls_i = set()
        ls_j = set()

        n = len(matrix)
        m = len(matrix[0])
        for i in range(n):
            for j in range(m):
                if matrix[i][j] == 0:
                    ls_i.add(i)
                    ls_j.add(j)

        for i in ls_i:
            for j in range(m):
                matrix[i][j] = 0
        
        for i in ls_j:
            for j in range(n):
                matrix[j][i] = 0
        
ls = [[1,1,1],[1,0,1],[1,1,1]]
res = Solution()
res.setZeroes(ls)
print(ls) 

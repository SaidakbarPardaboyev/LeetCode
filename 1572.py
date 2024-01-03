class Solution:
    def diagonalSum(self, mat: list[list[int]]) -> int:
        res = 0
        TheSame = 0

        for i in range(len(mat)):
            if len(mat[i]) > i:
                res += mat[i][i]

        j = len(mat) - 1
        for i in range(len(mat)):
            if len(mat[i]) > j:
                res += mat[i][j]
            j -= 1
        
        if len(mat) == len(mat[0]) and (len(mat) % 2 != 0 and len(mat[0]) % 2 != 0):
            tem = len(mat) // 2
            res -= mat[tem][tem]
        
        return res


print(Solution().diagonalSum([[1,2,3],
                              [4,5,6],
                              [7,8,9]]))

# 25

print(Solution().diagonalSum([[1,1,1,1],
                              [1,1,1,1],
                              [1,1,1,1],
                              [1,1,1,1]]))

# 8

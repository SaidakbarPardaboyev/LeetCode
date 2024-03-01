class Solution:
    def transpose(self, matrix: list[list[int]]) -> list[list[int]]:
        res = list()
        for _ in range(len(matrix[0])):
            res.append([])
        
        for i in range(len(matrix)):
            for j in range(len(matrix[0])):
                res[j].append(matrix[i][j])
        return res

print(Solution().transpose([[1,2,3],[4,5,6],[7,8,9]]))
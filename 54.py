class Solution:
    def spiralOrder(self, matrix: list[list[int]]) -> list[int]:
        res = list()

        lamp = 1
        while len(matrix) != 0:
            if lamp == 1:
                self.up(res, matrix)
                lamp += 1
            elif lamp == 2:
                self.right(res, matrix)
                lamp += 1
            elif lamp == 3:
                self.button(res, matrix)
                lamp += 1
            elif lamp == 4:
                self.left(res, matrix)
                lamp = 1
            j = 0
            while j < len(matrix):
                if not matrix[j]:
                    matrix.pop(j)
                else:
                    j += 1
        return res

    def up(self, res, matrix):
        res += matrix.pop(0)

    def right(self, res, matrix):
        leng = len(matrix)
        for i in range(leng):
            res.append(matrix[i][-1])
        for j in range(leng):
            matrix[j].pop()
    
    def button(self, res, matrix):
        res += matrix.pop()[::-1]

    def left(self, res, matrix):
        leng = len(matrix)
        for i in range(leng-1, -1, -1):
            res.append(matrix[i][0])
        for j in range(leng):
            matrix[j].pop(0)

print(Solution().spiralOrder([[1],[2],[3],[4],[5],[6],[7],[8],[9],[10]]))
print(Solution().spiralOrder([[7],[9],[6]]))
print(Solution().spiralOrder([[ 1, 2, 3, 4],
                              [ 5, 6, 7, 8],
                              [ 9,10,11,12],
                              [13,14,15,16]]))


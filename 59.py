class Solution:
    sch = 1
    def generateMatrix(self, n: int) -> list[list[int]]:
        matrix = list()
        res = list()
        com = [[None] * n for _ in range(n)]

        tem = list()
        for i in range(1, n*n+1):
            if i % n == 0:
                tem.append(i)
                matrix.append(tem.copy())
                res.append(tem.copy())
                tem = list()
            else:
                tem.append(i)

        lamp = 1
        while len(matrix) != 0:
            if lamp == 1:
                self.up(res, matrix, com)
                lamp += 1
            elif lamp == 2:
                self.right(res, matrix,com)
                lamp += 1
            elif lamp == 3:
                self.button(res, matrix,com)
                lamp += 1
            elif lamp == 4:
                self.left(res, matrix, com)
                lamp = 1
            j = 0
            while j < len(matrix):
                if not matrix[j]:
                    matrix.pop(j)
                else:
                    j += 1
        return com
    
    def up(self, res, matrix, com):
        leng = len(res)
        for i in range(len(matrix[0])):
            lamp = False
            for j in range(leng):
                for k in range(leng):
                    if res[j][k] == matrix[0][i]:
                        com[j][k] = self.sch
                        self.sch += 1
                        lamp = True
                        break
                if lamp:
                    break
        matrix.pop(0)

    def right(self, res, matrix, com):
        lengM = len(matrix)
        leng = len(res)
        for i in range(lengM):
            lamp = False
            for j in range(leng):
                for k in range(leng):
                    if res[j][k] == matrix[i][-1]:
                        com[j][k] = self.sch
                        self.sch += 1
                        lamp = True
                        break
                if lamp:
                    break
        for j in range(lengM):
            matrix[j].pop()

    def button(self, res, matrix, com):
        leng = len(res)
        for i in range(len(matrix[-1])-1, -1, -1):
            lamp = False
            for j in range(leng):
                for k in range(leng):
                    if res[j][k] == matrix[-1][i]:
                        com[j][k] = self.sch
                        self.sch += 1
                        lamp = True
                        break
                if lamp:
                    break
        matrix.pop()[::-1]

    def left(self, res, matrix, com):
        lengM = len(matrix)
        leng = len(res)
        for i in range(lengM-1, -1, -1):
            lamp = False
            for j in range(leng):
                for k in range(leng):
                    if res[j][k] == matrix[i][0]:
                        com[j][k] = self.sch
                        self.sch += 1
                        lamp = True
                        break
                if lamp:
                    break
        for j in range(lengM):
            matrix[j].pop(0)

print(Solution().generateMatrix(3))

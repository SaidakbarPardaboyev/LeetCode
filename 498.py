class Solution:
    def findDiagonalOrder(self, mat: list[list[int]]) -> list[int]:
        ls = []
        for k in range(0, len(mat[0])):
            if k % 2 == 0:
                j = 0
                i = k
                while i >= 0:
                    ls.append(mat[i][j])
                    i -= 1
                    j += 1
            else:
                i = 0
                j = k
                while j >= 0:
                    ls.append(mat[i][j])
                    i += 1
                    j -= 1

        for k in range(1, len(mat)):
            if k % 2 != 0:
                i = k
                j = len(mat)-1
                while j >= 1:
                    ls.append(mat[i][j])
                    i += 1
                    j -= 1
            else:
                i = len(mat)-1
                j = k
                while j < len(mat):
                    ls.append(mat[i][j])
                    i -= 1
                    j += 1
        return ls
    
print(Solution().findDiagonalOrder([[1, 2, 3],
                                    [4, 5, 6],
                                    [7, 8, 9]]))

            # [1,2,4,7,5,3,6,8,9]
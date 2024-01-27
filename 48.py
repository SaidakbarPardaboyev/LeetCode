from copy import deepcopy
class Solution:
    def rotate(self, matrix: list[list[int]]) -> None:
        """
        Do not return anything, modify matrix in-place instead.
        """
        tem = deepcopy(matrix)
        leng = len(matrix)
        for i in range(leng):
            for j in range(leng):
                matrix[j][-(i+1)] = tem[i][j]
                
        return matrix

print(Solution().rotate([[5,1,9,11],[2,4,8,10],[13,3,6,7],[15,14,12,16]]))
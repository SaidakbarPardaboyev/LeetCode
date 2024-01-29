class Solution:
    def searchMatrix(self, matrix: list[list[int]], target: int) -> bool:
        for i in range(len(matrix)):
            start = 0
            finish = len(matrix[i]) - 1

            while start <= finish:
                mid = (start + finish) // 2

                if matrix[i][mid] == target:
                    return True
                elif matrix[i][mid] > target:
                    finish = mid - 1
                else:
                    start =  mid + 1

        return False
    
print(Solution().searchMatrix([[1,3,5,7],[10,11,16,20],[23,30,34,60]], 3))
print(Solution().searchMatrix([[1,3,5,7],[10,11,16,20],[23,30,34,60]], 13))
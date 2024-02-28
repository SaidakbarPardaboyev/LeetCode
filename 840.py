class Solution:
    def numMagicSquaresInside(self, grid: list[list[int]]) -> int:
        res = 0
        ls = list()
        st = set()
        unique = {1,2,3,4,5,6,7,8,9}

        for k in range(len(grid)-2):
            for w in range(len(grid[0])-2):
                for i in range(k, k+3):
                    tem = list()
                    for j in range(w, w+3):
                        tem.append(grid[i][j])
                        st.add(grid[i][j])
                    ls.append(tem)

                if self.check(ls) and st == unique:
                    res += 1
                ls = list()
                st = set()

        return res

    def check(self, ls) -> bool:
        row1 = sum(ls[0])
        row2 = sum(ls[1])
        row3 = sum(ls[2])

        if row1 != row2 or row2 != row3:
            return False
        
        col1 = ls[0][0] + ls[1][0] + ls[2][0]
        col2 = ls[0][1] + ls[1][1] + ls[2][1]
        col3 = ls[0][2] + ls[1][2] + ls[2][2]

        if row1 != col1 or col1 != col2 or col2 != col3:
            return False

        dio1 = ls[0][0] + ls[1][1] + ls[2][2]
        dio2 = ls[0][2] + ls[1][1] + ls[2][0]

        if col1 != dio1 or dio1 != dio2:
            return False
        return True


print(Solution().numMagicSquaresInside([[2,7,6],[1,5,9],[4,3,8]]))
class Solution:
    def isValidSudoku(self, board: list[list[str]]) -> bool:
        for i in board:
            for j in i:
                if j != ".":
                    if i.count(j) > 1:
                        return False
        
        for i in range(9):
            tem = set()
            sch = 0
            for j in range(9):
                if board[j][i] != ".":
                    tem.add(board[j][i])
                    sch += 1
            if len(tem) != sch:
                return False

        box = set()
        sch = 0

        for i in range(3):
            for j in range(3):
                if board[i][j] != ".":
                    box.add(board[i][j])
                    sch += 1
        if len(box) != sch:
            return False

        
        box = set()
        sch = 0

        for i in range(3):
            for j in range(3, 6):
                if board[i][j] != ".":
                    box.add(board[i][j])
                    sch += 1
        if len(box) != sch:
            return False

        
        box = set()
        sch = 0

        for i in range(3):
            for j in range(6, 9):
                if board[i][j] != ".":
                    box.add(board[i][j])
                    sch += 1
        if len(box) != sch:
            return False

        box = set()
        sch = 0

        for i in range(3, 6):
            for j in range(3):
                if board[i][j] != ".":
                    box.add(board[i][j])
                    sch += 1
        if len(box) != sch:
            return False

        
        box = set()
        sch = 0

        for i in range(3,6):
            for j in range(3, 6):
                if board[i][j] != ".":
                    box.add(board[i][j])
                    sch += 1
        if len(box) != sch:
            return False

        
        box = set()
        sch = 0

        for i in range(3,6):
            for j in range(6, 9):
                if board[i][j] != ".":
                    box.add(board[i][j])
                    sch += 1
        if len(box) != sch:
            return False

        box = set()
        sch = 0

        for i in range(6,9):
            for j in range(3):
                if board[i][j] != ".":
                    box.add(board[i][j])
                    sch += 1
        if len(box) != sch:
            return False

        
        box = set()
        sch = 0

        for i in range(6,9):
            for j in range(3, 6):
                if board[i][j] != ".":
                    box.add(board[i][j])
                    sch += 1
        if len(box) != sch:
            return False

        
        box = set()
        sch = 0

        for i in range(6,9):
            for j in range(6, 9):
                if board[i][j] != ".":
                    box.add(board[i][j])
                    sch += 1
        if len(box) != sch:
            return False

        box = set()
        sch = 0
        return True
    
print(Solution().isValidSudoku([["5","3",".",".","7",".",".",".","."]
,["6",".",".","1","9","5",".",".","."]
,[".","9","8",".",".",".",".","6","."]
,["8",".",".",".","6",".",".",".","3"]
,["4",".",".","8",".","3",".",".","1"]
,["7",".",".",".","2",".",".",".","6"]
,[".","6",".",".",".",".","2","8","."]
,[".",".",".","4","1","9",".",".","5"]
,[".",".",".",".","8",".",".","7","9"]]))


print(Solution().isValidSudoku([["8","3",".",".","7",".",".",".","."]
,["6",".",".","1","9","5",".",".","."]
,[".","9","8",".",".",".",".","6","."]
,["8",".",".",".","6",".",".",".","3"]
,["4",".",".","8",".","3",".",".","1"]
,["7",".",".",".","2",".",".",".","6"]
,[".","6",".",".",".",".","2","8","."]
,[".",".",".","4","1","9",".",".","5"]
,[".",".",".",".","8",".",".","7","9"]]))
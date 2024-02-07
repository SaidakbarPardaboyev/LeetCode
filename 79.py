import os
os.system("cls")

class Solution:
    def exist(self, board: list[list[str]], word: str) -> bool:
        ls_begin_with = list()

        for i in range(len(board)):
            for j in range(len(board[0])):
                if board[i][j] == word[0]:
                    ls_begin_with.append([i, j])

        if not len(ls_begin_with):
            return False

        n = len(board)
        m = len(board[0])

        for pos in ls_begin_with:
            wd = word
            ls_past = list()
            list_way = list()
            tem = pos + [wd[1:]]
            list_way.append(tem)
            while True:
                cur_pos = list_way[-1]
                bor = 0
                if len(cur_pos[2]) == 0:
                    return True
                # move up 
                if cur_pos[0] - 1 >= 0:
                    if board[cur_pos[0]-1][cur_pos[1]] == cur_pos[2][0] and [cur_pos[0]-1, cur_pos[1]] not in ls_past:
                        bor += 1
                        list_way.append([cur_pos[0]-1] + [cur_pos[1]] + [cur_pos[2][1:]])

                # move right
                if cur_pos[1] + 1 < m:
                    if board[cur_pos[0]][cur_pos[1]+1] == cur_pos[2][0] and [cur_pos[0], cur_pos[1]+1] not in ls_past:
                        bor += 1
                        list_way.append([cur_pos[0]] + [cur_pos[1]+1] + [cur_pos[2][1:]])

                # move down
                if cur_pos[0] + 1 < n:
                    if board[cur_pos[0]+1][cur_pos[1]] == cur_pos[2][0] and [cur_pos[0]+1, cur_pos[1]] not in ls_past:
                        bor += 1
                        list_way.append([cur_pos[0]+1] + [cur_pos[1]] + [cur_pos[2][1:]])


                # move left
                if cur_pos[1] - 1 >= 0:
                    if (board[cur_pos[0]][cur_pos[1]-1] == cur_pos[2][0]) and ([cur_pos[0], cur_pos[1]-1] not in ls_past):
                        bor += 1
                        list_way.append([cur_pos[0]] + [cur_pos[1]-1] + [cur_pos[2][1:]])

                list_way.pop(list_way.index(cur_pos))
                if bor == 0:
                    break
                if len(list_way) == 0:
                    break
                ls_past.append([cur_pos[0], cur_pos[1]])

        return False        

print(Solution().exist([["A","B","C","E"],["S","F","E","S"],["A","D","E","E"]],
                       "ABC"))
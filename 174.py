from itertools import product

class Solution:
    def calculateMinimumHP(self, dungeon: list[list[int]]) -> int:
        if len(dungeon) == 1 and len(dungeon[0]) >= 2:
            turlar = ["right"]
        else:
            turlar = ["right", "down"]

        ls = product(turlar, repeat=len(dungeon[0]) + len(dungeon) - 2)
        # print(*ls)
        ls_right_com = list()
        for i in ls:
            if i.count("right") == len(dungeon[0])-1 or i.count("down") == len(dungeon)-1:
                ls_right_com.append(i)

        for i in ls_right_com:
            print(i)

        res = dungeon[0][0]
        for i in range(len(ls_right_com)):
            tem = res
            tem2 = tem
            for j in range(len(ls_right_com[i])):
                if ls_right_com[i][j] == "right":
                    tem += dungeon[i][j+1]
                    if tem < 0:
                        tem2 = min(tem, tem2)
                else:
                    tem += 
            res = max(res, tem2 * (-1))
        if res == dungeon[0][0]:
            return 1
        return res + 1
    
print(Solution().calculateMinimumHP([[0,-3]]))
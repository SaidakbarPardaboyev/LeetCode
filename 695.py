class Solution:
    def __init__(self):
        # Berilgan griddan hamma method larda foydalanish uchun
        # pointer kabi o'zgaruvchi ochich
        self.ls = None

    def maxAreaOfIsland(self, grid: list[list[int]]) -> int:
        # Berilgan griddan hamma method larda foydalanish uchun
        self.ls = grid

        # bitta listga orollarga tegishli kordinatalarni solib chiqish
        ls = []
        for i in range(len(grid)):
            for j in range(len(grid[i])):
                if grid[i][j] == 1:
                    ls.append([i, j])
        
        # Maximum qiymatni solish uchun o'zgaruvchi ochish
        Maxx = 0

        # Har bitta Kordinatadan boshlab yuramiz orolni hamma joyiga
        for i in ls:
            # Oldin bu joydan yurganmizmi tekshirish
            # Agar yurgan bo'lsak 1 ni 0 ga o'zgartirgan bo'lamiz
            if grid[i[0]][i[1]] == 0:
                continue
        
            # yo'l bo'linganda yurmagan yo'limizni kordinatalarini solish
            # uchun list ochish va birinchi turgan kordinatani solish
            cur_ls = list()
            cur_ls.append([i[0], i[1]])

            # yurayotgan orolomizni area sini hisoblash uchun o'zgaruvchi ochish
            area = 0

            while len(cur_ls) > 0:
                area += 1

                # joriy kordinata uchun o'zgaruvchi
                cur = cur_ls[-1]
                grid[cur[0]][cur[1]] = 0
                cur_ls.pop(-1)

                # oldimizda biz yura oladigan joylarni kordinatalarini olish uchun o'zgaruvchi
                tem = self.FindAllArea(cur[0], cur[1])
                if len(tem) > 0:
                    for i in tem:
                        if i not in cur_ls:
                            cur_ls.append(i)

            Maxx = max(Maxx, area)
    
        # for i in grid:
        #     print(i)
        
        # print()

        # for i in self.ls:
        #     print(i)
            
        return Maxx

    def FindAllArea(self, row, col) -> list[list[int]]:
        # Jo'natish uchun list ochish
        res = []

        # yuqoriga yura oladimi tekshirish
        if row > 0:
            if self.ls[row-1][col] == 1:
                res.append([row-1, col])
        
        # o'ngga yura oladimi tekshirish
        if col + 1 < len(self.ls[0]):
            if self.ls[row][col+1] == 1:
                res.append([row, col+1])
        
        # pastga yura oladimi tekshirish
        if row + 1 < len(self.ls):
            if self.ls[row+1][col] == 1:
                res.append([row+1, col])
        
        # Chapga yura oladimi tekshirish
        if col > 0:
            if self.ls[row][col-1] == 1:
                res.append([row, col-1])

        return res

print(Solution().maxAreaOfIsland([
    [0,0,1,0,0,0,0,1,0,0,0,0,0],
    [0,0,0,0,0,0,0,1,1,1,0,0,0],
    [0,1,1,0,1,0,0,0,0,0,0,0,0],
    [0,1,0,0,1,1,0,0,1,0,1,0,0],
    [0,1,0,0,1,1,0,0,1,1,1,0,0],
    [0,0,0,0,0,0,0,0,0,0,1,0,0],
    [0,0,0,0,0,0,0,1,1,1,0,0,0],
    [0,0,0,0,0,0,0,1,1,0,0,0,0]
    ]))

print(Solution().maxAreaOfIsland([
    [1,1,0,0,0],
    [1,1,0,0,0],
    [0,0,0,1,1],
    [0,0,0,1,1]
    ]))
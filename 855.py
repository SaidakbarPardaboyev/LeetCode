import os
os.system("cls")

class ExamRoom:

    def __init__(self, n: int):
        # N ta bo'sh o'rindiq yaratish
        self.ls = list()
        for _ in range(n):
            self.ls.append(0)

    def seat(self) -> int:
        # Odamlar bor joylar uchun list
        ls_indexs = list()

        # Odamlar bor joylatni topish
        for i in range(len(self.ls)):
            if self.ls[i] == 1:
                ls_indexs.append(i)
        
        if len(ls_indexs) == 0:
            self.ls[0] = 1
            return 0
        elif len(ls_indexs) == 1:
            if ls_indexs[0] < len(self.ls) / 2:
                self.ls[len(self.ls) - 1] = 1
                return len(self.ls) - 1
            else:
                self.ls[0] = 1
                return 0
        else:
            maxx = 0
            ind = 0
            if maxx < ls_indexs[0]:
                ind = 0
                maxx = ls_indexs[0]
            for i in range(len(ls_indexs)-1):
                tem = (ls_indexs[i+1] - ls_indexs[i]) // 2
                if maxx < tem:
                    maxx = tem
                    ind = ls_indexs[i] + tem
            tem = len(self.ls)-1 - ls_indexs[-1]
            if maxx < tem:
                maxx = tem
                ind = len(self.ls)-1
            self.ls[ind] = 1
            return ind

    def leave(self, p: int) -> None:
        self.ls[p] = 0  

# [null,0,9,4,null,null,0,4,2,6,1,3,5,7,8,null]
# Your ExamRoom object will be instantiated and called as such:
obj = ExamRoom(8)
param_1 = obj.seat()
print(param_1, end=" ")
param_1 = obj.seat()
print(param_1, end=" ")
param_1 = obj.seat()
print(param_1, end=" ")
print(obj.leave(0), end=" ")
print(obj.leave(7), end=" ")
param_1 = obj.seat()
print(param_1, end=" ")
param_1 = obj.seat()
print(param_1, end=" ")
param_1 = obj.seat()
print(param_1, end=" ")
param_1 = obj.seat()
print(param_1, end=" ")
param_1 = obj.seat()
print(param_1, end=" ")
param_1 = obj.seat()
print(param_1, end=" ")
param_1 = obj.seat()
print(param_1, end=" ")
param_1 = obj.seat()
print(param_1, end=" ")
# param_1 = obj.seat()
# print(param_1, end=" ")
# print(obj.leave(0), end=" ")


class Solution:
    def maxDistToClosest(self, seats: list[int]) -> int:
        # Odamlar bor joylar uchun list
        ls_indexs = list()

        # Odamlar bor joylatni topish
        for i in range(len(seats)):
            if seats[i] == 1:
                ls_indexs.append(i)
        
        # Bitta odam bo'lsa tekshiradi
        if len(ls_indexs) == 1:
            if ls_indexs[0] < len(seats) / 2:
                return len(seats) - ls_indexs[0] - 1
            else:
                return ls_indexs[0]
        # Ko'p odam bo'lsa tekshiradi
        else:
            maxx = 0
            for i in range(len(ls_indexs)-1):
                maxx = max(maxx, (ls_indexs[i+1] - ls_indexs[i]) // 2)
            # Boshidagi ba oxiridagi bo'sh joylarni tekshiradi
            maxx = max(maxx, ls_indexs[0])
            maxx = max(maxx, len(seats)-1 - ls_indexs[-1])
            return maxx

print(Solution().maxDistToClosest([1,0,0,0,0,0,1,0,1]))
print(Solution().maxDistToClosest([1,0,1]))
print(Solution().maxDistToClosest([1,0,0,0]))

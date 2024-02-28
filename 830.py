class Solution:
    def largeGroupPositions(self, s: str) -> list[list[int]]:
        res = list()
        # solishtirilayotgan belgilarni Boshlang'ich index
        start = 0
        # solishtirilayotgan belgilarni Oxirgi index
        end = 0
        # Solishtirish uchun belgini saqlab turish uchun
        belgi = s[0]
        # Har bitta harfni tekshirib chiqadi
        for i in range(1, len(s)):
            # Belgilar o'xshash bo'lsa i index ni end ga solib ketadi 
            if s[i] == belgi:
                end = i
                # Oxirgi belgini ham list ga qo'shish uchun tekshirish
                if i == len(s)-1:
                    if end-start >= 2:
                        res.append([start, end])    
            else:
                # Listga qo'shish uchun shart
                if end-start >= 2:
                    res.append([start, end])
                start = i
                belgi = s[i]
        
        return res
    
print(Solution().largeGroupPositions( "abbxxxxzzy"))
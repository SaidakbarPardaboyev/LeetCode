class Solution:
    def halvesAreAlike(self, s: str) -> bool:
        ls = ['a', 'e', 'i', 'o', 'u']
        half1 = s[:len(s)//2]
        half2 = s[len(s)//2:]

        sch1 = 0
        sch2 = 0

        for i in range(len(half1)):
            if half1[i].lower() in ls:
                sch1 += 1
            if half2[i].lower() in ls:
                sch2 += 1
        
        return sch1 == sch2


print(Solution().halvesAreAlike("textbook"))
print(Solution().halvesAreAlike("book"))

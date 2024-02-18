class Solution:
    def convertToBase7(self, num: int) -> str:
        if num == 0:
            return "0"
        summa = ""
        lamp = False
        if num < 0:
            num = -num
            lamp = True
        while num > 0:
            summa += str(num % 7)
            num //= 7
        
        if lamp:
            return '-' + summa[::-1] 
        return summa[::-1]

print(Solution().convertToBase7(0))

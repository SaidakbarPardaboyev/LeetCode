class Solution:
    def multiply(self, num1: str, num2: str) -> str:
        num1_int = 0
        num2_int = 0

        for i in num1:
            num1_int = num1_int * 10 + (ord(i)-ord("0"))
        
        for i in num2:
            num2_int = num2_int * 10 + (ord(i)-ord("0"))
        
        tem = num1_int * num2_int

        if tem == 0:
            return "0"

        res = ""

        while tem != 0:
            n = tem % 10
            res += chr(ord("0")+n)
            tem = tem // 10
        
        return res[::-1]
    
print(Solution().multiply("123", "456"))
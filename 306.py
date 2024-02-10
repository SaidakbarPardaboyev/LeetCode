class Solution:
    def isAdditiveNumber(self, num: str) -> bool:
        for i in range(1, len(num)):
            for j in range(1, len(num)):
                if len(num[:i]) == len(str(int(num[:i]))) and len(num[i:i+j]) == len(str(int(num[i:i+j]))):
                    if num[i+j:].startswith(str(int(num[:i])+int(num[i:i+j]))):
                        if self.Check(num[:i], num[i:i+j], num[i+j:]):
                            return True
                        else:
                            continue
        return False
    
    def Check(self, a, b, num) -> bool:
        while len(num):
            c = str(int(a) + int(b))
            if num.startswith(c):
                a = b
                b = c
                num = num[len(c):]
            else:
                return False
        return True
    
print(Solution().isAdditiveNumber("123"))
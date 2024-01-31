class Solution:
    def fractionToDecimal(self, numerator: int, denominator: int) -> str:
        res = numerator / denominator
        if res == int(res):
            return str(int(res))
        else:
            return self.checkFraction(str(res))
            
    def checkFraction(self, string: str) -> str:
        indDot = string.find('.')
        res = string[:indDot+1]
        fraction = string[indDot + 1:]
        i = 0
        while i < len(fraction) // 2:
            for j in range(len(fraction)):
                tem = (fraction[j:j+i + 1] * ((len(fraction) - i) // len(fraction[j:j+i + 1]) - 1))
                if fraction[j:-1] == tem and len(tem) != 1:
                    res += (fraction[:j] + f"({fraction[j:j+i + 1]})")
                    return res
            i+=1
        
        return str(string)

# print(Solution().fractionToDecimal(4, 333))
print(Solution().fractionToDecimal(14, 8))
print(Solution().fractionToDecimal(10, 3))


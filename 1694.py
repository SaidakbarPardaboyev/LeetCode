class Solution:
    def reformatNumber(self, number: str) -> str:
        res = str()
        number = number.replace('-', '')
        number = number.replace(' ', '')

        while len(number) > 4:
            res += (number[:3] + '-')
            number = number[3:]
        
        if len(number) == 4:
            res += (number[:2] + '-' + number[2:])
        else:
            res += number
        return res
        

print(Solution().reformatNumber("1-23-45 6"))
print(Solution().reformatNumber("123 4-567"))
print(Solution().reformatNumber("123 4-5678"))

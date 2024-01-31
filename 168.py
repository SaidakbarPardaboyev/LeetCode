class Solution:
    def convertToTitle(self, columnNumber: int) -> str:
        if not columnNumber > 0:
            return ""
        
        columnNumber -= 1
        qoldiq = columnNumber % 26
        columnNumber = columnNumber // 26
        return self.convertToTitle(columnNumber) + chr(ord('A') + qoldiq)
        
print(Solution().convertToTitle(701))
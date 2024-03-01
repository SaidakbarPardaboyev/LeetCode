class Solution:
    def binaryGap(self, n: int) -> int:
        binar = bin(n)[2:]
        if not binar.count('1') > 1:
            return 0
        
        res = 1
        count = 1
        lamp = False

        for char in binar:
            if char == '1':
                if lamp == False:
                    lamp = True
                else:
                    res = max(res, count)
                    count = 1
            else:
                count += 1
        return res

print(Solution().binaryGap(8))
class Solution:
    def decrypt(self, code: list[int], k: int) -> list[int]:
        lamp = False
        if k == 0:
            for i in range(len(code)):
                code[i] = 0
            return code        
        elif k < 0:
            k *= -1
            code.reverse()
            lamp = True
        
        ls = list()
        size = len(code)
        for i in range(size):
            tem = 0
            for j in range(1, k+1):
                if i + j >= size:
                    tem += code[(i + j) - size]
                else:
                    tem += code[i + j]
            ls.append(tem)

        if lamp:
            ls.reverse()
        return ls

print(Solution().decrypt([2,4,9,3], -2))
                        #   0 1 2 3             
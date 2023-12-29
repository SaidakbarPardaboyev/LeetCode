class Solution:
    def finalPrices(self, prices: list[int]) -> list[int]:
        res = list()
        
        for i in range(len(prices) - 1):
            tem = prices[0]
            prices.pop(0)
            FirstMin = 0
            for j in prices:
                if j <= tem:
                    FirstMin = j
                    break
            tem -= FirstMin 
            res.append(tem)

        res.append(prices[0])
        return res
            
result = Solution()
print(result.finalPrices([8,4,6,2,3]))
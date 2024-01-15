class Solution:
    def maximumWealth(self, accounts: list[list[int]]) -> int:
        Max = sum(accounts[0])

        for i in range(1, len(accounts)):
            Max = max(Max, sum(accounts[i]))
        
        return Max
    
print(Solution().maximumWealth([[1,5],[7,3],[3,5]]))
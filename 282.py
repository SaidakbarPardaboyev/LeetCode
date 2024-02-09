class Solution:
    def addOperators(self, num: str, target: int) -> list[str]:
        from itertools import product
        items = ["", '+', '*', '-']
        ishora = ['+', '*', '-']
        ls_com = product(items, repeat=len(num)-1)
        res = list()
        for i in ls_com:
            tem = num[0]
            for j in range(len(i)):
                tem += (i[j] + num[j+1])
            lamb = True
            for k in range(1,len(tem)):
                if tem[k] == '0':
                    if tem[k-1] in ishora:
                        if len(tem) > k+1:
                            if not tem[k+1] in ishora:
                                lamb = False
                    elif tem[k-1] == '0':
                        lamb = False
            if lamb:
                if eval(tem) == target:
                    res.append(tem)
        return res
    
res = Solution().addOperators("00", 0)

for i in res:
    print(i)
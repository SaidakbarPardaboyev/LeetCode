class Solution:
    def lemonadeChange(self, bills: list[int]) -> bool:
        dic = {5:0, 10:0, 20:0}

        for i in bills:
            if i == 10:
                if dic[5] > 0:
                    dic[i] += 1
                    dic[5] -= 1
                else:
                    return False
                continue

            if i == 20:
                if dic[10] > 0:
                    if dic[5] > 0:
                        dic[10] -= 1
                        dic[5] -= 1
                    else:
                        return False
                elif dic[5] > 2:
                    dic[5] -= 3
                else:
                    return False
                dic[20] += 1
                continue
                
            dic[i] += 1
        return True
    
print(Solution().lemonadeChange([5,5,5,10,20]))
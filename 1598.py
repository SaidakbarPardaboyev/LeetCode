class Solution:
    def minOperations(self, logs: list[str]) -> int:
        ls = []
        for i in logs:
            if i.endswith("../"):
                if len(ls) != 0:
                    ls.pop()
            elif i.endswith("/") and not i.endswith("./"):
                ls.append(i[:-1])
        return len(ls)

print(Solution().minOperations(["d1/","d2/","../","d21/","./"]))
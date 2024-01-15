class Solution:
    def countConsistentStrings(self, allowed: str, words: list[str]) -> int:
        allowed = set(allowed)
        sch = 0

        for i in words:
            ls = list(i)
            lamp = False
            for j in ls:
                if j not in allowed:
                    lamp = True
                    break
            if not lamp:
                sch += 1
        return sch
    
print(Solution().countConsistentStrings("ab", ["ad","bd","aaab","baa","badab"]))
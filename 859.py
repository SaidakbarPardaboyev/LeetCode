class Solution:
    def buddyStrings(self, s: str, goal: str) -> bool:
        if len(s) != len(goal):
            return False

        if s == goal:
            if len(set(s)) != len(s):
                return True
            else:
                return False

        Sbelgi1 = None
        Sbelgi2 = None
        Gbelgi1 = None
        Gbelgi2 = None
        count = 0

        for i in range(len(s)):
            if s[i] != goal[i]:
                if count == 0:
                    Sbelgi1 = s[i]
                    Gbelgi1 = goal[i]
                elif count == 1:
                    Sbelgi2 = s[i]
                    Gbelgi2 = goal[i]
                else:
                    return False
                count += 1
        
        if Sbelgi1 == Gbelgi2 and Sbelgi2 == Gbelgi1:
            return True
        else:
            return False

print(Solution().buddyStrings("ab", "ba"))
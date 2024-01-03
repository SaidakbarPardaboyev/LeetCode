class Solution:
    def containsPattern(self, arr: list[int], m: int, k: int) -> bool:
        i = 0
        while i + m < len(arr):
            count = 1
            FirstMElement = arr[i:i + m]
            j = i + m
            while j + m <= len(arr):
                SecondMElement = arr[j:j + m]
                if SecondMElement == FirstMElement:
                    count += 1
                else:
                    break
                j += 1
            if count == k:
                return True
            i += 1

        return False

print(Solution().containsPattern([2,2,2,2], 2, 3))

        # 1 2 4
        # 2 4 4 
        # 4 4 4
        # 4 4 4
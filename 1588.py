class Solution:
    def sumOddLengthSubarrays(self, arr: list[int]) -> int:
        summa = sum(arr)
        i = 3
        size = len(arr)
        while i <= size:
            k = 0
            while k + i <= size:
                tem = 0
                for j in range(k, k + i):
                    tem += arr[j]
                summa += tem
                k += 1
            i += 2
        return summa

print(Solution().sumOddLengthSubarrays([1,4,2,5,3]))
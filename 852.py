class Solution:
    def peakIndexInMountainArray(self, arr: List[int]) -> int:
        # return arr.index(max(arr))
        maxx = arr[1]
        ind = 1
        i = 1
        j = len(arr)-1
        while i <= j:
            if arr[i] > maxx:
                maxx = arr[i]
                ind = i
            if arr[j] > maxx:
                maxx = arr[j]
                ind = j
            i+=1
            j-=1
        
        return ind
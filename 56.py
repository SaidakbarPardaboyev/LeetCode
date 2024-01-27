class Solution:
    def merge(self, intervals: list[list[int]]) -> list[list[int]]:
        intervals.sort(key=lambda x: x[0])
        i = 0
        while i < len(intervals)-1:
            if intervals[i][-1] >= intervals[i+1][-1]:
                intervals.pop(i+1)
            elif intervals[i][-1] >= intervals[i+1][0]:
                intervals[i] = [intervals[i][0], intervals[i+1][-1]]
                intervals.pop(i+1)
            else:
                i += 1
        return intervals
    
print(Solution().merge([[1,3],[2,6],[8,10],[15,18]]))
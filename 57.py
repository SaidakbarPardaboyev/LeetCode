class Solution:
    def insert(self, intervals: list[list[int]], newInterval: list[int]) -> list[list[int]]:
        intervals.append(newInterval)
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


print(Solution().insert([[1,2],[3,5],[6,7],[8,10],[12,16]], [4,8]))
print(Solution().insert([[1,3],[6,9]], [2,5]))
print(Solution().insert([[1,4],[6,9]], [2,4]))
print(Solution().insert([[1,2],[3,5],[6,7],[8,15],[19,21]], [4,15]))

# print(Solution().insert([[1,3],[6,9]], [2,5]))

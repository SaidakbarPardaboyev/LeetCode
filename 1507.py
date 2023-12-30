class Solution:
    def reformatDate(self, date: str) -> str:
        month = {"Jan": 1, "Feb": 2, "Mar": 3, "Apr": 4, "May": 5, "Jun": 6, "Jul": 7, "Aug": 8, "Sep": 9, "Oct": 10, "Nov": 11, "Dec": 12}

        date = date.split(' ')

        return f"{date[-1]}-{str(month[date[1]]).zfill(2)}-{date[0][:-2].zfill(2)}"
    
print(Solution().reformatDate("6th Jun 1933"))
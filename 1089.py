def duplicateZeros(arr: list[int]) -> None:
        """
        Do not return anything, modify arr in-place instead.
        """
        i = 1
        while i < len(arr):
            if arr[i - 1] == 0:
                arr.insert(i, 0)
                arr.pop()
                i += 1
            i += 1

print(duplicateZeros([1,0,2,3,0,4,5,0]))

        # [1,0,0,2,3,0,0,4]
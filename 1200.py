def minimumAbsDifference(arr: list[int]) -> list[list[int]]:
    arr.sort()
    min_diff = float('inf')
    result = []

    for i in range(len(arr) - 1):
        diff = arr[i + 1] - arr[i]
        min_diff = min(min_diff, diff)

    for i in range(len(arr) - 1):
        if arr[i + 1] - arr[i] == min_diff:
            result.append([arr[i], arr[i + 1]])

    return result

print(minimumAbsDifference([40,11,26,27,-20]))
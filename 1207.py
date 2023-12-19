def checkStraightLine(coordinates: list[list[int]]) -> bool:
    if len(coordinates) <= 2:
        return True

    if coordinates[1][0] == coordinates[0][0]:
        x_val = coordinates[0][0]
        for i in range(2, len(coordinates)):
            if coordinates[i][0] != x_val:
                return False
        return True

    slope = (coordinates[1][1] - coordinates[0][1]) / (coordinates[1][0] - coordinates[0][0])

    for i in range(2, len(coordinates)):
        if coordinates[i][0] - coordinates[0][0] == 0:
            return False

        current_slope = (coordinates[i][1] - coordinates[0][1]) / (coordinates[i][0] - coordinates[0][0])
        if current_slope != slope:
            return False

    return True

print(checkStraightLine([[1,2],[2,3],[3,4],[4,5],[5,6],[6,7]]))
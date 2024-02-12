func furthestDistanceFromOrigin(moves string) int {
    res := 0
    schL := 0
    schR := 0
    for _, i := range moves {
        if i == 'L' {
            schL += 1
        } else if i == 'R' {
            schR += 1
        }
    }
    count := 0
    if schL > schR {
        count = -1
    } else {
        count = 1
    }

    for _, i := range moves {
        if i == 'L' {
            res -= 1
        } else if i == 'R' {
            res += 1
        } else {
            res += count
        }
    }
    if res < 0 {
        return -res
    }
    return res
}
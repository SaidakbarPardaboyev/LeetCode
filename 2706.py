def buyChoco(prices: list[int], money: int) -> int:
    min1 = min(prices)
    prices.remove(min1)
    min2 = min(prices)
    if min1 + min2 > money:
        return money
    else:
        return money - (min1 + min2)
    
print(buyChoco([1,2,2], 3))
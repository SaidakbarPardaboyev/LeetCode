func maxBottlesDrunk(numBottles int, numExchange int) int {
	res := 0
	for numBottles-numExchange >= 0 {
		numBottles = numBottles - numExchange + 1
		res += numExchange
		numExchange += 1
	}
	res += numBottles
	return res
}
import "math"

func findTheCity(n int, edges [][]int, distanceThreshold int) int {
	citiesWithNeighbors := findCitiesNeighbors(edges)

	smallestreachableCitesNumber := n + 1
	cityWithSmallreachableCites := -1
	for i := 0; i < n; i++ {
		costToReachCites := findCostToReachCites(citiesWithNeighbors, i)
		reachableCitesNumber := 0
		for _, cost := range costToReachCites {
			if cost <= distanceThreshold {
				reachableCitesNumber++
			}
		}
		if smallestreachableCitesNumber > reachableCitesNumber {
			smallestreachableCitesNumber = reachableCitesNumber
			cityWithSmallreachableCites = i
		} else if smallestreachableCitesNumber == reachableCitesNumber {
			if i > cityWithSmallreachableCites {
				smallestreachableCitesNumber = reachableCitesNumber
				cityWithSmallreachableCites = i
			}
		}
	}

	return cityWithSmallreachableCites
}

func findCostToReachCites(citiesWithNeighbors map[int]map[int]int, start int) map[int]int {

	costsToGoCities := map[int]int{start: 0}
	fromTo := map[int]int{}
	visited := map[int]bool{}

	cheapestCity := start
	for cheapestCity != -1 {
		costToReachCurrentCity := costsToGoCities[cheapestCity]
		neighbors := citiesWithNeighbors[cheapestCity]

		for neighbor, cost := range neighbors {
			newCost := costToReachCurrentCity + cost
			if _, ok := costsToGoCities[neighbor]; !ok || costsToGoCities[neighbor] > newCost {
				costsToGoCities[neighbor] = newCost
				fromTo[neighbor] = cheapestCity
			}
		}

		visited[cheapestCity] = true
		cheapestCity = findCheapestCity(costsToGoCities, visited)
	}
	return costsToGoCities
}

func findCheapestCity(costs map[int]int, visited map[int]bool) int {
	cheapestCost := math.MaxInt64
	cheapestCity := -1

	for city, cost := range costs {
		if cost < cheapestCost && !visited[city] {
			cheapestCost = cost
			cheapestCity = city
		}
	}
	return cheapestCity
}

func findCitiesNeighbors(edges [][]int) map[int]map[int]int {
	citiesWithNeighbors := make(map[int]map[int]int)
	for _, edge := range edges {
		from := edge[0]
		to := edge[1]
		price := edge[2]

		if _, ok := citiesWithNeighbors[from]; !ok {
			citiesWithNeighbors[from] = map[int]int{}
		}

		if _, ok := citiesWithNeighbors[to]; !ok {
			citiesWithNeighbors[to] = map[int]int{}
		}

		citiesWithNeighbors[from][to] = price
		citiesWithNeighbors[to][from] = price
	}
	return citiesWithNeighbors
}
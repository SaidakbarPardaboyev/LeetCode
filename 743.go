package main

import (
	"fmt"
	"math"
)

func main(){
	edges := [][]int{{2,1,1},{2,3,1},{3,4,1}}
	n := 4
	k := 2

	fmt.Println(networkDelayTime(edges, n, k))
	fmt.Println(networkDelayTime([][]int{{1,2,1}, {2,1,3}}, 2, 2))
}

func networkDelayTime(edges [][]int, n int, k int) int {
	pathsAmongCities := FindAllCitiesNetworks(edges)

	parents := map[int]int{}
	times := map[int]int{k: 0}
	visited := map[int]bool{}

	currentNode := k
	for currentNode != -1 {
		curruntTime := times[currentNode]
		neighbours := pathsAmongCities[currentNode]
		for neighbour, time := range neighbours{
			newTimes := time + curruntTime
			if _, ok := times[neighbour]; !ok || times[neighbour] > newTimes {
				times[neighbour] = newTimes
				parents[neighbour] = currentNode
			}
		}
		visited[currentNode] = true
		currentNode = FindSmallTimeOfADestination(times, visited)
	}

    if _, ok := times[n]; !ok {
        return -1
    }

    if _, ok := times[k]; !ok {
        return -1
    }

	for key := range pathsAmongCities {
		if _, ok := times[key]; !ok {
			return -1
		}
	}

	maxTime := 0
    for _, time := range times {
        if time > maxTime {
            maxTime = time
        }
    }
    fmt.Println(times)
    return maxTime
}

func FindSmallTimeOfADestination(times map[int]int, visited map[int]bool) int {
	cityWithSmallTime := -1
	smallTimeToReach := math.MaxInt64

	for city, time := range times {
		if !visited[city] && time < smallTimeToReach {
			smallTimeToReach = time
			cityWithSmallTime = city
		}
	}
	return cityWithSmallTime
}

func FindAllCitiesNetworks(paths [][]int) map[int]map[int]int {
	pathsAmongCities := make(map[int]map[int]int)

	for _, path := range paths {
		from := path[0]
		to := path[1]
		time := path[2]
		
		if _, ok := pathsAmongCities[from]; !ok {
			pathsAmongCities[from] = map[int]int{}
		}

		pathsAmongCities[from][to] = time
	}

	return pathsAmongCities
}
package main

import (
	"fmt"
	"math"
)

// Vertex for a weighted graph, name is the name of the actual vertex (like a city name)
// Edges is a map of the edges and the cost to get to that vertex.
type Vertex struct {
	name  string
	edges map[*Vertex]int
}

func (v *Vertex) addAdjacent(name *Vertex, cost int) {
	v.edges[name] = cost
}

func newCity(name string) Vertex {
	return Vertex{
		name:  name,
		edges: map[*Vertex]int{},
	}
}

func main() {
	atlanta := newCity("Atlanta")
	boston := newCity("Boston")
	chicago := newCity("Chicago")
	denver := newCity("Denver")
	elPaso := newCity("El Paso")

	atlanta.addAdjacent(&boston, 100)
	atlanta.addAdjacent(&denver, 160)
	boston.addAdjacent(&chicago, 120)
	boston.addAdjacent(&denver, 180)
	chicago.addAdjacent(&elPaso, 80)
	denver.addAdjacent(&chicago, 40)
	denver.addAdjacent(&elPaso, 140)

	fmt.Println(findShortest(atlanta, "El Paso"))
}

// Dijkstras shortest path
// Steps
// 1.	Make the starting city the current city
// 2.	Check the prices from the current city to each of its adjacent cities
// 3.	If the price is cheaper than the value saved in currentLowestPrices or unset
//		a.	Update the currentLowestPrices table to reflect this new cheaper price
//		b.	Update the cheapestPreviousStopoverCity, making the adjacent city the key
//			and the current city the value
// 4.	Then visit whichever unvisited city has the cheapest price from the starting city,
//		making it the current city
// 5.	Repeat step 2-4 until every known city has been visited

// Find the shortest path between a vertex in a graph and the endPoint, which will be
// the name (string) of a Vertex type.
// I think it should return the sum of the values in the path, and the path itself.
func findShortest(startingPoint Vertex, endPoint string) string {
	lowestCost := map[string]int{}
	lowestPrevious := map[string]string{}

	unvisitedVertices := []*Vertex{&startingPoint}
	visitedVertices := map[string]struct{}{}

	lowestCost[startingPoint.name] = 0
	currentVertex := startingPoint

	for len(unvisitedVertices) != 0 {
		visitedVertices[currentVertex.name] = struct{}{}
		unvisitedVertices = removeCurrent(unvisitedVertices, currentVertex.name)

		for vertex, cost := range currentVertex.edges {
			if _, ok := visitedVertices[vertex.name]; !ok {
				unvisitedVertices = append(unvisitedVertices, vertex)
			}
			currentCost := lowestCost[currentVertex.name] + cost

			if _, ok := lowestCost[vertex.name]; !ok {
				lowestCost[vertex.name] = currentCost
				lowestPrevious[vertex.name] = currentVertex.name
			} else if lowestCost[vertex.name] > currentCost {
				lowestCost[vertex.name] = currentCost
				lowestPrevious[vertex.name] = currentVertex.name
			}

		}

		currentVertex = minUnvisited(unvisitedVertices, lowestCost)
	}

	path := []string{endPoint}
	curr := lowestPrevious[endPoint]
	for curr != startingPoint.name {
		path = append(path, curr)
		curr = lowestPrevious[curr]
	}
	path = append(path, startingPoint.name)

	result := ""

	for i := len(path) - 1; i > 0; i-- {
		result += fmt.Sprintf("%s => ", path[i])
	}
	result += path[0]

	return fmt.Sprintf("The shortest path is: %s\nThe cost is %d", result, lowestCost[endPoint])
}

func minUnvisited(notSeen []*Vertex, lowestCost map[string]int) Vertex {
	var res Vertex
	currLowCost := math.MaxInt

	for _, vertex := range notSeen {
		if lowestCost[vertex.name] < currLowCost {
			res = *vertex
			currLowCost = lowestCost[vertex.name]
		}
	}

	return res
}

func removeCurrent(vertices []*Vertex, name string) []*Vertex {
	res := []*Vertex{}
	for i, v := range vertices {
		if v.name == name {
			res := vertices[:i]
			res = append(res, vertices[i+1:]...)
			break
		}
	}

	return res
}

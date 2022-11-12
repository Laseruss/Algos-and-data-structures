package main

import "fmt"

// Struct type for each vertex in a graph
type Vertex struct {
	value    interface{}
	adjacent []*Vertex
}

// Constructor that returns a new vertex with value set to passed argument
// and an empty slice for adjacent vertices.
func constructor(value interface{}) Vertex {
	return Vertex{
		value:    value,
		adjacent: []*Vertex{},
	}
}

// Adds a vertex as adjacent to the vertex that called the method, also sets
// the vertex called as adjacent on the vertex that is passed in. If the vertex is
// already adjacent nothing happens.
func (v *Vertex) add(val *Vertex) {
	for _, vertex := range v.adjacent {
		if vertex == val {
			return
		}
	}

	v.adjacent = append(v.adjacent, val)
	val.adjacent = append(val.adjacent, v)
}

// Performs depth first search through the graph, for now it prints out the value
// of every vertex in the graph. Needs to be passed a map to know which vertices
// that has already been visited.
func (v *Vertex) depthFirstTraverse(visited map[*Vertex]bool) {
	visited[v] = true
	fmt.Println(v.value)

	for _, vertex := range v.adjacent {
		if _, ok := visited[vertex]; !ok {
			vertex.depthFirstTraverse(visited)
		}
	}
}

func (v *Vertex) depthFirstSearch(want interface{}, visited map[*Vertex]bool) *Vertex {
	if v.value == want {
		return v
	}
	visited[v] = true

	for _, vertex := range v.adjacent {
		if _, ok := visited[vertex]; ok {
			continue
		}

		if vertex.value == want {
			return vertex
		}

		searchedVertex := vertex.depthFirstSearch(want, visited)

		if searchedVertex != nil {
			return searchedVertex
		}
	}

	return nil
}

func (v *Vertex) BreadthFirstTraverse(visited map[interface{}]bool) {
	visited[v.value] = true
	queue := []*Vertex{v}

	for len(queue) > 0 {
		q, curr := dequeue(queue)
		queue = q

		fmt.Println(curr.value)

		for _, v := range curr.adjacent {
			if _, ok := visited[v.value]; !ok {
				visited[v.value] = true
				queue = enqueue(v, queue)
			}
		}
	}
}

func enqueue(value *Vertex, l []*Vertex) []*Vertex {
	l = append(l, value)
	return l
}

func dequeue(l []*Vertex) ([]*Vertex, *Vertex) {
	return l[1:], l[0]
}

func main() {
	alice := constructor("alice")
	bob := constructor("bob")
	cynthia := constructor("cynthia")

	alice.add(&bob)
	alice.add(&cynthia)
	bob.add(&cynthia)
	cynthia.add(&bob)

	cynthia.depthFirstTraverse(map[*Vertex]bool{})

	fmt.Println(alice.depthFirstSearch("glenn", map[*Vertex]bool{}))

	alice.BreadthFirstTraverse((map[interface{}]bool{}))
}

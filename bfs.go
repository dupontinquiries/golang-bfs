package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/gammazero/deque"
)

// makes a randomized adjacency matrix
func m2d(s int, mirror bool) [][]bool {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	m := make([][]bool, s)
	for i := 0; i < s; i++ {
		m[i] = make([]bool, s)
	}
	for i := 0; i < s*s; i++ {
		j := int(i / s)
		// generating one-by-one is inefficient & should be replaced by https://stackoverflow.com/questions/45030618/generate-a-random-bool-in-go
		m[i-(s*j)][j] = r1.Intn(2+int(s/4)) == 0
		// mirror for undirected graph
		if mirror {
			m[j][i-(s*j)] = m[i-(s*j)][j]
		}
	}
	return m
}

func sm(m [][]bool, show_axes bool) string {
	var r string
	var s int = len(m)
	for i := -1; i < s; i++ {
		for j := -1; j < s; j++ {
			r += " "
			if j == -1 && i == -1 {
				if show_axes {
					r += "  "
				}
			} else if i == -1 && j != -1 {
				if show_axes {
					r += fmt.Sprint(j) + "."
				}
				// r += fmt.Sprint(j) + "."
			} else if j == -1 && i != -1 {
				if show_axes {
					r += fmt.Sprint(i) + "."
				}
				// r += fmt.Sprint(i) + "."
			} else if i == j {
				r += "X "
			} else if m[i][j] {
				r += "1 "
			} else if i != -1 && j != -1 {
				r += "0 "
			}
		}
		r += "\n"
	}
	return r
}

// TODO: add visual node display (show circle representing each node & show each edge as a line connecting two nodes)

func bfs(m [][]bool, start int, end int) int {
	// TODO: clean up code and try to implement path memory to show all nodes in shortest path
	fmt.Printf("starting bfs for %d -> %d\n", start, end)
	// queue of visited nodes
	var v []bool = make([]bool, len(m))
	// distance int
	var d int
	// queue of nodes (size, min capacity)
	q := deque.New(len(m)*len(m), len(m))
	path := deque.New(len(m), len(m))
	// add starting int(node) to queue
	q.PushFront(start)
	path.PushBack(start)
	// add starting node to visited
	v[start] = true
	// while queue not empty
	for q.Len() != 0 {
		// ++distance
		d++
		// work on first node in queue
		i := int(q.PopFront().(int))
		// debugging
		fmt.Println("addressing node: ", i)
		// add all true values in m[i](neighbors of node i) to the queue
		for j := 0; j < len(m[i]); j++ {
			// k := 0
			if i != j && !v[j] && m[i][j] {
				if j == end {
					// path.PushBack(j)
					fmt.Printf("found end node! (d=%d)\n", d)
					// path_slice := make([]int, path.Len())
					// for tmp := 0; path.Len() != 0; tmp++ {
					// path_slice[tmp] = int(path.PopFront().(int))
					// }
					// fmt.Println("path:", path_slice)
					return d
				}
				fmt.Println(" - found neighbor: ", j)
				q.PushBack(j)
				v[j] = true
				// k++
				// path.PushBack(j)
			}
			// for l := 0; l < k-1; l++ {
			// path.PopBack()
			// }
		}
	}
	fmt.Println("did not find end node")
	return -1
}

func main() {
	// make a 2d adjacency matrix using pseudo-random values
	//  - an adjacency matrix is a matrix that denotes edges as ones
	//  - a zero means that two nodes are not directly connected by an edge
	var n int = 8
	var is_bidirectional bool = true
	// get start and end values
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	start := r1.Intn(n)
	end := start
	for start == end {
		end = r1.Intn(n)
	}
	fmt.Println("goal:", start, "->", end)
	// n: number of nodes
	// is_bidirectional: whether edges are bidirectional or directed
	m := m2d(n, is_bidirectional)
	// show the adjacency matrix to help visualize what is going on
	var show_axes bool = n <= 10
	fmt.Println("matrix representation:\n" + sm(m, show_axes))
	// fmt.Println("graph representation:\n" + sg(m))
	bfs(m, start, end)
}

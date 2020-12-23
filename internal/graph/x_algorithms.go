package graph

import (
	"container/heap"

	"github.com/ngeorgiadis/baseline/internal/dataset"
)

// FindLBack ...
func (G XGraph) FindLBack(authors *dataset.AuthorPriorityQueue, L int) []XGraph {

	res := []XGraph{}

	for {
		top := heap.Pop(authors).(*dataset.Author)
		n := G[top.ID].Clone()

		for n.Visited {
			top = heap.Pop(authors).(*dataset.Author)
			n = G[top.ID].Clone()
		}

		n.Visited = true
		n.IsInitialCommunityNode = true

		community := XGraph{}
		community.AddNode(n)

		for to := range n.Edges {

			v := G[to]
			for to2 := range v.Edges {
				w := G[to2]
				if w.HasNeighbour(n) {
					v.Visited = true
					community.AddNode(v)
					break
				}
			}
		}

		res = append(res, community)

		if len(res) >= L {
			break
		}
	}

	return res
}

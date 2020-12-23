package graph

import (
	"container/heap"
	"sort"

	"github.com/ngeorgiadis/baseline/internal/dataset"
)

// FindTopDomCommunity ...
func (G IxGraph) FindTopDomCommunity(authors *dataset.AuthorPriorityQueue) IxGraph {

	top := heap.Pop(authors).(*dataset.Author)

	n := G[top.ID]
	n.Visited = true

	community := IxGraph{}
	community.AddNode(n)

	for {

		sort.SliceStable(n.Neighbours, func(i, j int) bool {
			return n.Neighbours[i].Data.GetDomScore() > n.Neighbours[j].Data.GetDomScore()
		})

		n = n.Neighbours.First(func(item *Node) bool {
			return !item.Visited
		})

		if n == nil {
			break
		}

		n.Visited = true
		community.AddNode(n)
	}

	return community
}

// Find1Back ...
func (G IxGraph) Find1Back(authors *dataset.AuthorPriorityQueue) IxGraph {

	top := heap.Pop(authors).(*dataset.Author)
	n := G[top.ID]
	n.Visited = true

	community := IxGraph{}
	community.AddNode(n)

	for _, v := range n.Neighbours {
		for _, w := range v.Neighbours {
			if w.HasNeighbour(n) {
				community.AddNode(v)
				break
			}
		}
	}

	return community
}

// FindLBack ...
func (G IxGraph) FindLBack(authors *dataset.AuthorPriorityQueue, L int) []IxGraph {

	res := []IxGraph{}

	for {
		top := heap.Pop(authors).(*dataset.Author)
		n := G[top.ID]

		for n.Visited {
			top = heap.Pop(authors).(*dataset.Author)
			n = G[top.ID]
		}

		n.Visited = true
		n.IsInitialCommunityNode = true

		community := IxGraph{}
		community.AddNode(n)
		for _, v := range n.Neighbours {

			for _, w := range v.Neighbours {

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

// FindLBack2 ...
func (G IxGraph) FindLBack2(authors *dataset.AuthorPriorityQueue, L int) []IxGraph {

	res := []IxGraph{}

	for {
		top := heap.Pop(authors).(*dataset.Author)
		n := G[top.ID]

		for n.Visited {
			top = heap.Pop(authors).(*dataset.Author)
			n = G[top.ID]
		}

		n.Visited = true
		n.IsInitialCommunityNode = true

		community := IxGraph{}
		community.AddNode(n)
		for _, v := range n.Neighbours {

			//for _, w := range v.Neighbours {

			ns := v.GetAllNeighboursAdjTo(n)
			community.AddNodes(ns)

			// if w.HasNeighbour(n) {
			// 	v.Visited = true
			// 	community.AddNode(v)
			// 	break
			// }
			//}
		}

		res = append(res, community)

		if len(res) >= L {
			break
		}
	}

	return res
}

// MaxKCore ...
func (G IxGraph) MaxKCore() {

}

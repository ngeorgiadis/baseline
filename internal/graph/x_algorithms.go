package graph

import (
	"container/heap"

	"github.com/ngeorgiadis/baseline/internal/dataset"
)

// FindLBack ...
func (G XGraph) FindLBack(authors *dataset.AuthorPriorityQueue, L int) []XGraph {

	res := []XGraph{}
	visited := map[string]bool{}

	for {
		top := heap.Pop(authors).(*dataset.Author)

		n := G[top.ID].Clone()
		for visited[n.ID] {
			top = heap.Pop(authors).(*dataset.Author)
			n = G[top.ID].Clone()
		}

		visited[n.ID] = true
		n.IsInitialCommunityNode = true

		community := XGraph{}
		community.AddNode(n)

		for edge := range n.Edges {

			v := G[edge].Clone()
			for vEdge := range v.Edges {

				w := G[vEdge]
				if w.HasNeighbour(n) {

					visited[v.ID] = true
					community.AddNode(v)
					break
				}

			}
		}

		// clean nil edges
		idx := map[string]struct{}{}
		for k := range community {
			idx[k] = struct{}{}
		}

		for _, c := range community {
			for r := range c.Edges {
				if _, ok := idx[r]; !ok {
					delete(c.Edges, r)
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

// KCore ...
func (G XGraph) KCore(k int) XGraph {

	for _, n := range G {
		n.Visited = false
	}

	for _, n := range G {
		G.DFS(k, n)
	}

	return G
}

// DFS ...
func (G XGraph) DFS(k int, n *XNode) {

	if n.Visited {
		return
	}

	n.Visited = true

	if len(n.Edges) < k {

		for i := range n.Edges {

			adj := G[i]
			if adj == nil {
				delete(n.Edges, i)
			} else {
				delete(adj.Edges, n.ID)
				adj.Visited = false
			}

			// if adj, ok := G[i]; ok {
			// delete(adj.Edges, n.ID)
			// adj.Visited = false
			// }
		}

		delete(G, n.ID)

		for i := range n.Edges {

			adj := G[i]
			if adj == nil {
				delete(n.Edges, i)
			} else {
				G.DFS(k, adj)
			}

			// if adj, ok := G[i]; ok {
			// 	G.DFS(k, adj)
			// }
		}

	}
}

// MaxKCore ...
func (G XGraph) MaxKCore() (XGraph, int) {

	i := 0

	var maxG XGraph
	// var kG XGraph

	for {

		kG := G.KCore(i)

		if (len(kG)) <= 0 {
			break
		}

		i++
		maxG = kG.Clone()
	}

	return maxG, i - 1

}

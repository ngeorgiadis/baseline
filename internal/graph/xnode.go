package graph

import "github.com/ngeorgiadis/baseline/internal/dataset"

// XNode ...
type XNode struct {
	ID                     string
	Edges                  map[string]struct{}
	Visited                bool
	Data                   DominanceReaderWriter
	IsInitialCommunityNode bool
	index                  int
}

// AddEdge ...
func (n *XNode) AddEdge(edgeTo string) {
	n.Edges[edgeTo] = struct{}{}
}

// HasNeighbour ...
func (n *XNode) HasNeighbour(v *XNode) bool {
	res := false
	for to := range n.Edges {
		if v.ID == to {
			return true
		}
	}
	return res
}

// Clone ...
func (n *XNode) Clone() *XNode {

	data2 := *n.Data.(*dataset.Author)

	res := &XNode{
		ID:                     n.ID,
		Visited:                n.Visited,
		IsInitialCommunityNode: n.IsInitialCommunityNode,
		Edges:                  map[string]struct{}{},
		index:                  n.index,
		Data:                   &data2,
	}

	for e := range n.Edges {
		res.Edges[e] = struct{}{}
	}

	return res
}

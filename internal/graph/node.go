package graph

import (
	"container/heap"
	"sort"
)

// DominanceReaderWriter ...
type DominanceReaderWriter interface {
	GetDomScore() int
	SetDomScore(v int)
}

// Node ...
type Node struct {
	ID                     string
	Neighbours             SortedNodes
	Visited                bool
	Data                   DominanceReaderWriter
	IsInitialCommunityNode bool

	index int
}

// NewNode ...
func NewNode() *Node {
	n := &Node{}
	return n
}

// AddNeighbour ...
func (n *Node) AddNeighbour(a *Node) {

	n.Neighbours = append(n.Neighbours, a)

	sort.SliceStable(n.Neighbours, func(i, j int) bool {
		return n.Neighbours[i].Data.GetDomScore() > n.Neighbours[j].Data.GetDomScore()
	})
	// n.Neighbours = append(n.Neighbours, a)
	// heap.Push(&n.Neighbours, a)
}

// HasNeighbour ...
func (n *Node) HasNeighbour(v *Node) bool {
	res := false
	for _, i := range n.Neighbours {
		if v.ID == i.ID {
			return true
		}
	}
	return res
}

// GetAllNeighboursAdjTo ...
func (n *Node) GetAllNeighboursAdjTo(v *Node) []*Node {
	res := []*Node{}
	for _, i := range n.Neighbours {
		for _, j := range i.Neighbours {
			if v.ID == j.ID {
				res = append(res, i)
			}
		}
	}
	return res
}

// SortedNodes ...
type SortedNodes []*Node

func (sn SortedNodes) Len() int { return len(sn) }

func (sn SortedNodes) Less(i, j int) bool {
	return sn[i].Data.GetDomScore() > sn[j].Data.GetDomScore()
}

func (sn SortedNodes) Swap(i, j int) {
	sn[i], sn[j] = sn[j], sn[i]
	sn[i].index = i
	sn[j].index = j
}

// First ...
func (sn SortedNodes) First(p func(*Node) bool) *Node {
	for _, n := range sn {
		if p(n) {
			return n
		}
	}

	return nil
}

// NodesPriorityQueue ...
type NodesPriorityQueue []*Node

func (npq NodesPriorityQueue) Len() int { return len(npq) }

func (npq NodesPriorityQueue) Less(i, j int) bool {
	return npq[i].Data.GetDomScore() > npq[j].Data.GetDomScore()
}

func (npq NodesPriorityQueue) Swap(i, j int) {
	npq[i], npq[j] = npq[j], npq[i]
	npq[i].index = i
	npq[j].index = j
}

// Push ...
func (npq *NodesPriorityQueue) Push(x interface{}) {
	n := len(*npq)
	item := x.(*Node)
	item.index = n
	*npq = append(*npq, item)
}

// Pop ...
func (npq *NodesPriorityQueue) Pop() interface{} {
	old := *npq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*npq = old[0 : n-1]
	return item
}

// Update ...
func (npq *NodesPriorityQueue) Update(item *Node, score int) {
	item.Data.SetDomScore(score)
	heap.Fix(npq, item.index)
}

// First ...
func (npq NodesPriorityQueue) First(p func(*Node) bool) *Node {
	for _, n := range npq {
		if p(n) {
			return n
		}
	}

	return nil
}

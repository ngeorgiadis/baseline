package graph

import (
	"container/heap"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/ngeorgiadis/baseline/internal/dataset"
)

// Graph ...
type Graph []*Node

// IxGraph ...
type IxGraph map[string]*Node

// XGraph ...
type XGraph map[string]*XNode

// New ...
func New() Graph {
	return Graph{}
}

// NewIx ...
func NewIx() IxGraph {
	return IxGraph{}
}

// NewXGraph ...
func NewXGraph() XGraph {
	return XGraph{}
}

// AddNode ...
func (g Graph) AddNode(n *Node) {
	g = append(g, n)
}

// AddNode ...
func (g IxGraph) AddNode(n *Node) {
	g[n.ID] = n
}

// AddNodes ...
func (g IxGraph) AddNodes(nodes []*Node) {
	for _, n := range nodes {
		g[n.ID] = n
	}
}

// AddNode ...
func (g XGraph) AddNode(n *XNode) {
	g[n.ID] = n
}

// AddNodes ...
func (g XGraph) AddNodes(nodes []*XNode) {
	for _, n := range nodes {
		g[n.ID] = n
	}
}

// Clone ...
func (g XGraph) Clone() XGraph {

	newXG := XGraph{}

	for _, n := range g {
		nn := *n
		nn.Edges = map[string]struct{}{}
		for e := range n.Edges {
			nn.Edges[e] = struct{}{}
		}
		newXG.AddNode(&nn)
	}

	return newXG
}

// FromFile ...
func FromFile(nodesFile string, edgesFile string) (XGraph, dataset.AuthorPriorityQueue) {

	// read input data
	// and construct indexes
	f, _ := os.Open(nodesFile)
	r := csv.NewReader(f)

	// nodes (authors)
	nodeRecs, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	// edged (authors edges)
	f, _ = os.Open(edgesFile)
	r = csv.NewReader(f)
	edgeRecs, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	xG := XGraph{}
	authors := dataset.AuthorPriorityQueue{}

	// NODES
	for i, row := range nodeRecs {
		if i == 0 {
			continue
		}

		id := row[0]
		name := row[1]
		pn, _ := strconv.Atoi(row[2])
		cn, _ := strconv.Atoi(row[3])
		hi, _ := strconv.ParseFloat(row[4], 64)
		pi, _ := strconv.ParseFloat(row[5], 64)

		a := &dataset.Author{
			ID:                id,
			Name:              name,
			PublicationNumber: pn,
			CitationNumber:    cn,
			HIndex:            hi,
			PIndex:            pi,
		}

		authors = append(authors, a)

		// add to index
		n := &XNode{
			ID:      id,
			Data:    a,
			Visited: false,
			Edges:   map[string]struct{}{},
		}

		xG.AddNode(n)
	}

	// EDGES
	for i, row := range edgeRecs {
		if i == 0 {
			continue
		}
		sourceID := row[0]
		targetID := row[1]

		s := xG[sourceID]
		t := xG[targetID]

		s.AddEdge(t.ID)
		t.AddEdge(s.ID)

	}

	// create authors
	// Prioriry queue
	l := len(authors)
	domMap := map[string][]*dataset.Author{}
	for i := 0; i < l; i++ {
		for j := 0; j < l; j++ {

			if i == j {
				continue
			}

			a := authors[i]
			b := authors[j]

			if a.Dominates(b) {
				a.DomScore++

				if _, ok := domMap[a.ID]; !ok {
					domMap[a.ID] = []*dataset.Author{}
				}
				t := domMap[a.ID]
				t = append(t, b)
				domMap[a.ID] = t
			}
		}
	}

	heap.Init(&authors)

	return xG, authors
}

// ToFile ...
func (g XGraph) ToFile(name string) {

	f, err := os.Create(name)
	if err != nil {
		fmt.Printf(err.Error())
	}

	defer f.Close()

	for _, n := range g {

		a := n.Data.(*dataset.Author)

		edges := ""
		for e := range n.Edges {
			edges += e + "|"
		}

		fmt.Fprintf(f, "%v,%v,%v,%v\n", n.ID, edges, n.Visited, a.Name)
	}

}

package main

// import (
// 	"encoding/csv"
// 	"fmt"
// 	"log"
// 	"os"
// 	"sort"
// )

// // Print ...
// func (g Graph) Print() {
// 	for _, n := range g {
// 		adj := ""
// 		for _, a := range n.Neighbours {
// 			adj += a.NodeData.ID + ", "
// 		}
// 		fmt.Printf("%v -> %v", n.NodeData.ID, adj)
// 	}
// }

// // Neighbours ...
// type Neighbours []*Node

// func (g Graph) findTopKDomCommunities(index map[string]*Node, domMap map[string][]*Author, authors []*Author, k int) [][]Node {

// 	res := [][]Node{}

// 	// authors is sorted by dom score
// 	i := 0

// 	var top *Author
// 	top = authors[i]

// 	for {
// 		i++
// 		// the community
// 		comm := []Node{}

// 		// start here
// 		// index is the graph index
// 		// with ID as key
// 		n := index[top.ID]
// 		n.Visited = true
// 		top.Visited = true

// 		comm = append(comm, *n)

// 		for {

// 			neighbours := []*Node{}
// 			for _, i := range n.Neighbours {
// 				if !i.Visited {
// 					neighbours = append(neighbours, i)
// 				}
// 			}

// 			if len(neighbours) > 0 {
// 				sort.SliceStable(neighbours, func(i, j int) bool {
// 					return neighbours[i].NodeData.DomScore > neighbours[j].NodeData.DomScore
// 				})

// 				n = neighbours[0]
// 				n.NodeData.Visited = true

// 			} else {
// 				n = nil
// 			}

// 			if n == nil {
// 				break
// 			}
// 			n.Visited = true
// 			comm = append(comm, *n)
// 		}

// 		// create graph files
// 		// nodes and edges
// 		// from the node set (comm)
// 		// GenerateSubgraphFiles(comm, i)
// 		res = append(res, comm)

// 		// check if k it sat
// 		if i > k {
// 			break
// 		}

// 		// select new top
// 		// top = the next author with highest dom scroce
// 		// that is not visited
// 		var next *Author
// 		for _, t := range authors {
// 			if !t.Visited {
// 				next = t
// 				break
// 			}
// 		}

// 		// if top == nil break
// 		if next == nil {
// 			break
// 		}

// 		top = next

// 	}

// 	return res

// }

// // GenerateSubgraphFiles ...
// func GenerateSubgraphFiles(comm []Node, i int) {

// 	index := map[string]Node{}
// 	for _, n := range comm {
// 		index[n.ID] = n
// 	}

// 	f, _ := os.Open("./data/nodes_all.csv")
// 	r := csv.NewReader(f)
// 	records, err := r.ReadAll()
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	of, _ := os.Create(fmt.Sprintf("./data/nodes_sub%v_%v.csv", i, runID))
// 	writer := csv.NewWriter(of)

// 	for i, row := range records {
// 		if i == 0 {
// 			row = append(row, "domscore")
// 			writer.Write(row)
// 			continue
// 		}
// 		if n, ok := index[row[0]]; ok {
// 			row = append(row, fmt.Sprintf("%v", n.NodeData.DomScore))
// 			writer.Write(row)
// 		}
// 	}

// 	writer.Flush()
// 	of.Close()

// 	f, _ = os.Open("./data/edges.csv")
// 	r = csv.NewReader(f)
// 	records, err = r.ReadAll()
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	of, _ = os.Create(fmt.Sprintf("./data/edges_sub%v_%v.csv", i, runID))
// 	defer of.Close()
// 	writer = csv.NewWriter(of)

// 	for i, row := range records {
// 		if i == 0 {
// 			writer.Write(row)
// 			continue
// 		}

// 		_, ok := index[row[0]]
// 		_, ok2 := index[row[1]]

// 		if ok && ok2 {
// 			writer.Write(row)
// 		}
// 	}

// 	writer.Flush()

// }

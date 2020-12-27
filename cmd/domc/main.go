package main

import (
	"container/heap"
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/ngeorgiadis/baseline/internal/dataset"
	"github.com/ngeorgiadis/baseline/internal/graph"
)

var runID = ""

func main2() {

	runID = fmt.Sprintf("%v", time.Now().UnixNano())

	// read input data
	// and construct indexes
	f, _ := os.Open("./data/nodes.csv")
	r := csv.NewReader(f)

	// nodes (authors)
	nodeRecs, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	// edged (authors edges)
	f, _ = os.Open("./data/edges.csv")
	r = csv.NewReader(f)
	edgeRecs, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	//
	//
	ixG := graph.IxGraph{}

	//
	// *NODES*
	//
	authors := dataset.AuthorPriorityQueue{}
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
		n := &graph.Node{
			ID:         id,
			Data:       a,
			Neighbours: []*graph.Node{},
			Visited:    false,
		}

		ixG.AddNode(n)

		// add to graph
		// ????
	}

	//
	// *EDGES*
	//
	for i, row := range edgeRecs {
		if i == 0 {
			continue
		}
		sourceID := row[0]
		targetID := row[1]

		s := ixG[sourceID]
		t := ixG[targetID]

		s.AddNeighbour(t)
	}

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

	// for authors.Len() > 0 {
	// 	item := heap.Pop(&authors).(*dataset.Author)
	// 	fmt.Printf("%.2d:%s \n", item.DomScore, item.Name)
	// }

	// calculate top-k ...
	// ...

	// white nodes to file
	// just once to include domscore
	// writeNodes(ixG)

	// c1 := ixG.FindTopDomCommunity(&authors)
	// c2 := c1.FindTopDomCommunity(&authors)

	// fmt.Println(len(c1))
	// fmt.Println(len(c2))

	//
	// clean graph from
	//
	lc := ixG.FindLBack2(&authors, 1)

	// for i, c := range lc {
	// 	writeNodes(c, i)
	// }

	topL := graph.IxGraph{}

	for _, c := range lc {
		for _, n := range c {
			topL.AddNode(n)
		}
	}

}

func main() {

	runID = fmt.Sprintf("%v", time.Now().UnixNano())

	// read input data
	// and construct indexes
	f, _ := os.Open("./data/nodes.csv")
	r := csv.NewReader(f)

	// nodes (authors)
	nodeRecs, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	// edged (authors edges)
	f, _ = os.Open("./data/edges.csv")
	r = csv.NewReader(f)
	edgeRecs, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	//
	//
	xG := graph.XGraph{}

	//
	// *NODES*
	//
	authors := dataset.AuthorPriorityQueue{}
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
		n := &graph.XNode{
			ID:      id,
			Data:    a,
			Visited: false,
			Edges:   map[string]struct{}{},
		}

		xG.AddNode(n)

		// add to graph
		// ????
	}

	//
	// *EDGES*
	//
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

	// for authors.Len() > 0 {
	// 	item := heap.Pop(&authors).(*dataset.Author)
	// 	fmt.Printf("%.2d:%s \n", item.DomScore, item.Name)
	// }

	// calculate top-k ...
	// ...

	// white nodes to file
	// just once to include domscore
	// writeNodes(ixG)

	// c1 := ixG.FindTopDomCommunity(&authors)
	// c2 := c1.FindTopDomCommunity(&authors)

	// fmt.Println(len(c1))
	// fmt.Println(len(c2))

	//
	// clean graph from
	//
	lc := xG.FindLBack(&authors, 25)

	// remove edges that are not in
	// the graph

	idx := map[string]struct{}{}
	for _, community := range lc {
		for k := range community {
			idx[k] = struct{}{}
		}
	}

	for _, community := range lc {
		for _, c := range community {
			for r := range c.Edges {
				if _, ok := idx[r]; !ok {
					delete(c.Edges, r)
				}
			}
		}
	}

	// write
	// read
	f, _ = os.Open("./data/nodes_all.csv")
	r = csv.NewReader(f)
	allnodes, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	f, _ = os.Open("./data/edges_all.csv")
	r = csv.NewReader(f)
	alledges, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	for i, c := range lc {

		writeNodes2(c, i, allnodes, alledges)

		cMaxCore, m := c.MaxKCore()
		writeNodes3(cMaxCore, i, m, allnodes, alledges)
	}

	// // prune
	// for i := 0; i < len(lc); i++ {
	// 	lc[i], _ = lc[i].MaxKCore()
	// }

}

func writeNodes(ixG graph.IxGraph, i int) {

	os.MkdirAll(fmt.Sprintf("./data/%v/", runID), 0777)

	// read
	f, _ := os.Open("./data/nodes_all.csv")
	r := csv.NewReader(f)
	records, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	of, _ := os.Create(fmt.Sprintf("./data/%v/nodes_i%v.csv", runID, i))
	writer := csv.NewWriter(of)

	for i, row := range records {
		if i == 0 {
			row = append(row, "domscore")
			row = append(row, "degree")
			row = append(row, "is_initial")
			writer.Write(row)
			continue
		}

		if n, ok := ixG[row[0]]; ok {
			row = append(row, fmt.Sprintf("%v", n.Data.GetDomScore()))
			row = append(row, fmt.Sprintf("%v", len(n.Neighbours)))
			row = append(row, fmt.Sprintf("%v", n.IsInitialCommunityNode))
			writer.Write(row)
		}
	}

	writer.Flush()
	of.Close()

	f, _ = os.Open("./data/edges_all.csv")
	r = csv.NewReader(f)
	records, err = r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	of, _ = os.Create(fmt.Sprintf("./data/%v/edges_i%v.csv", runID, i))
	defer of.Close()
	writer = csv.NewWriter(of)

	for i, row := range records {
		if i == 0 {
			writer.Write(row)
			continue
		}

		_, ok := ixG[row[0]]
		_, ok2 := ixG[row[1]]

		if ok && ok2 {
			writer.Write(row)
		}
	}

	writer.Flush()
}

func writeNodes2(xG graph.XGraph, i int, allnodes [][]string, alledges [][]string) {

	os.MkdirAll(fmt.Sprintf("./data/%v/", runID), 0777)
	os.MkdirAll(fmt.Sprintf("../../data/ui/"), 0777)

	of, _ := os.Create(fmt.Sprintf("./data/%v/nodes_i%v.csv", runID, i))

	writer := csv.NewWriter(of)

	for i, row := range allnodes {
		if i == 0 {
			row = append(row, "domscore")
			row = append(row, "degree")
			row = append(row, "is_initial")
			writer.Write(row)
			continue
		}

		if n, ok := xG[row[0]]; ok {
			row = append(row, fmt.Sprintf("%v", n.Data.GetDomScore()))
			row = append(row, fmt.Sprintf("%v", len(n.Edges)))
			row = append(row, fmt.Sprintf("%v", n.IsInitialCommunityNode))
			writer.Write(row)
		}
	}

	writer.Flush()
	of.Close()

	b, _ := ioutil.ReadFile(fmt.Sprintf("./data/%v/nodes_i%v.csv", runID, i))
	ioutil.WriteFile(fmt.Sprintf("../../data/ui/nodes_i%v.csv", i), b, 0777)

	of, _ = os.Create(fmt.Sprintf("./data/%v/edges_i%v.csv", runID, i))

	defer of.Close()
	writer = csv.NewWriter(of)

	for i, row := range alledges {
		if i == 0 {
			writer.Write(row)
			continue
		}

		_, ok := xG[row[0]]
		_, ok2 := xG[row[1]]

		if ok && ok2 {
			writer.Write(row)
		}
	}

	writer.Flush()
	of.Close()

	b, _ = ioutil.ReadFile(fmt.Sprintf("./data/%v/edges_i%v.csv", runID, i))
	ioutil.WriteFile(fmt.Sprintf("../../data/ui/edges_i%v.csv", i), b, 0777)

}

func writeNodes3(xG graph.XGraph, i int, m int, allnodes [][]string, alledges [][]string) {

	os.MkdirAll(fmt.Sprintf("./data/%v/", runID), 0777)
	os.MkdirAll(fmt.Sprintf("../../data/ui/"), 0777)

	of, _ := os.Create(fmt.Sprintf("./data/%v/nodes_maxcore_i%v.csv", runID, i))

	writer := csv.NewWriter(of)

	for i, row := range allnodes {
		if i == 0 {
			row = append(row, "domscore")
			row = append(row, "degree")
			row = append(row, "is_initial")
			row = append(row, "max_k_core")

			writer.Write(row)
			continue
		}

		if n, ok := xG[row[0]]; ok {
			row = append(row, fmt.Sprintf("%v", n.Data.GetDomScore()))
			row = append(row, fmt.Sprintf("%v", len(n.Edges)))
			row = append(row, fmt.Sprintf("%v", n.IsInitialCommunityNode))
			row = append(row, fmt.Sprintf("%v", m))
			writer.Write(row)
		}
	}

	writer.Flush()
	of.Close()

	b, _ := ioutil.ReadFile(fmt.Sprintf("./data/%v/nodes_maxcore_i%v.csv", runID, i))
	ioutil.WriteFile(fmt.Sprintf("../../data/ui/nodes_maxcore_i%v.csv", i), b, 0777)

	of, _ = os.Create(fmt.Sprintf("./data/%v/edges_maxcore_i%v.csv", runID, i))

	defer of.Close()
	writer = csv.NewWriter(of)

	for i, row := range alledges {
		if i == 0 {
			writer.Write(row)
			continue
		}

		_, ok := xG[row[0]]
		_, ok2 := xG[row[1]]

		if ok && ok2 {
			writer.Write(row)
		}
	}

	writer.Flush()
	of.Close()

	b, _ = ioutil.ReadFile(fmt.Sprintf("./data/%v/edges_maxcore_i%v.csv", runID, i))
	ioutil.WriteFile(fmt.Sprintf("../../data/ui/edges_maxcore_i%v.csv", i), b, 0777)

}

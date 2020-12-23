package main

var runID = ""

func main() {

	// runID = fmt.Sprintf("%v", time.Now().UnixNano())

	// g := graph.Graph{}
	// f, _ := os.Open("./data/nodes.csv")
	// r := csv.NewReader(f)
	// records, err := r.ReadAll()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// authors := []*dataset.Author{}
	// index := map[string]*graph.Node{}

	// for i, row := range records {

	// 	if i == 0 {
	// 		continue
	// 	}

	// 	id := row[0]
	// 	name := row[1]
	// 	pn, _ := strconv.Atoi(row[2])
	// 	cn, _ := strconv.Atoi(row[3])
	// 	hi, _ := strconv.ParseFloat(row[4], 64)
	// 	pi, _ := strconv.ParseFloat(row[5], 64)

	// 	a := dataset.Author{
	// 		ID:                id,
	// 		Name:              name,
	// 		PublicationNumber: pn,
	// 		CitationNumber:    cn,
	// 		HIndex:            hi,
	// 		PIndex:            pi,
	// 	}

	// 	authors = append(authors, &a)
	// 	n := &graph.Node{
	// 		ID:         id,
	// 		Data:       &a,
	// 		Neighbours: []*graph.Node{},
	// 		Visited:    false,
	// 	}

	// 	g = append(g, n)

	// 	index[id] = n
	// }

	// // edges
	// f, _ = os.Open("./data/edges.csv")
	// r = csv.NewReader(f)
	// records, err = r.ReadAll()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// for i, row := range records {
	// 	if i == 0 {
	// 		continue
	// 	}
	// 	sourceID := row[0]
	// 	targetID := row[1]

	// 	s := index[sourceID]
	// 	t := index[targetID]

	// 	s.AddNeighbour(t)
	// }

	// // START

	// l := len(authors)
	// domMap := map[string][]*dataset.Author{}

	// for i := 0; i < l; i++ {
	// 	for j := 0; j < l; j++ {

	// 		if i == j {
	// 			continue
	// 		}

	// 		a := authors[i]
	// 		b := authors[j]

	// 		if a.Dominates(b) {
	// 			a.DomScore++

	// 			if _, ok := domMap[a.ID]; !ok {
	// 				domMap[a.ID] = []*Author{}
	// 			}
	// 			t := domMap[a.ID]
	// 			t = append(t, b)
	// 			domMap[a.ID] = t
	// 		}
	// 	}
	// }

	// sort.SliceStable(authors, func(i, j int) bool {
	// 	return authors[i].DomScore > authors[j].DomScore
	// })

	// comms := g.findTopKDomCommunities(index, domMap, authors, 3)

	// fmt.Println(len(comms))

	// o, _ := os.Create("out.csv")
	// o2, _ := os.Create("dom_index.csv")

	// defer o.Close()
	// defer o2.Close()

	// for _, a := range authors {
	// 	doms := domMap[a.ID]

	// 	fmt.Fprintf(o, "%v,%v,%v\n", a.ID, a.Name, a.DomScore)

	// 	if len(doms) > 0 {
	// 		fmt.Fprintf(o2, "%v,", a.ID)
	// 	} else {
	// 		fmt.Fprintf(o2, "%v\n", a.ID)
	// 	}

	// 	for i, d := range doms {
	// 		if i == len(doms)-1 {
	// 			fmt.Fprintf(o2, "%v\n", d.ID)
	// 		} else {
	// 			fmt.Fprintf(o2, "%v,", d.ID)
	// 		}
	// 	}
	// }

	// fmt.Printf("%v\n", mat.Formatted(dmat))

}

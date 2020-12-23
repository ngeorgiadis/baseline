package handlers

import (
	"baseline/m/v2/cmd/server/config"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"path"
	"strconv"
)

type graphViewModel struct {
	Nodes []graphNode `json:"nodes"`
	Edges []graphEdge `json:"edges"`
}

type graphNode struct {
	ID         string            `json:"id"`
	Label      string            `json:"label"`
	Attributes map[string]string `json:"attrs"`
	Cluster    string            `json:"cluster"`
	Shared     bool              `json:"shared"`
	IsInit     bool              `json:"is_init"`
	Degree     int               `json:"degree"`
}

type graphEdge struct {
	Source string `json:"source"`
	Target string `json:"target"`
}

func getGraphData(c *config.Config) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		res := graphViewModel{
			Nodes: []graphNode{},
			Edges: []graphEdge{},
		}

		dataFolder := "./data"

		nodesIndex := map[string]graphNode{}
		edgesIndex := map[string]graphEdge{}

		for fi := 0; fi < 10; fi++ {

			edgesFile := fmt.Sprintf("edges_1606422527546140500_i%v.csv", fi)
			nodesFile := fmt.Sprintf("nodes_1606422527546140500_i%v.csv", fi)

			f, _ := os.Open(path.Join(dataFolder, nodesFile))
			reader := csv.NewReader(f)
			recs, _ := reader.ReadAll()
			for i, rec := range recs {
				if i == 0 {
					continue
				}

				attr := map[string]string{
					"pc":       rec[2],
					"cn":       rec[3],
					"hi":       rec[4],
					"pi":       rec[5],
					"domscore": rec[6],
				}

				gn := graphNode{
					ID:         rec[0],
					Label:      rec[1],
					Attributes: attr,
					Cluster:    fmt.Sprintf("%v", fi),
					Shared:     false,
				}

				if _, exists := nodesIndex[gn.ID]; exists {
					gn.Shared = true
				}
				nodesIndex[gn.ID] = gn

				// res.Nodes = append(res.Nodes, graphNode{
				// 	ID:         rec[0],
				// 	Label:      rec[1],
				// 	Attributes: attr,
				// 	Cluster:    fmt.Sprintf("%v", i),
				// })
			}

			f, _ = os.Open(path.Join(dataFolder, edgesFile))
			reader = csv.NewReader(f)
			recs, _ = reader.ReadAll()
			for i, rec := range recs {
				if i == 0 {
					continue
				}

				edgesIndex[fmt.Sprintf("%v%v", rec[0], rec[1])] = graphEdge{
					Source: rec[0],
					Target: rec[1],
				}

				// res.Edges = append(res.Edges, graphEdge{
				// 	Source: rec[0],
				// 	Target: rec[1],
				// })
			}

		}

		for _, n := range nodesIndex {
			res.Nodes = append(res.Nodes, n)
		}

		for _, e := range edgesIndex {
			res.Edges = append(res.Edges, e)
		}

		b, _ := json.Marshal(&res)
		w.Write(b)
	}
}

// 1606980136635444400
func getGraphData2(c *config.Config) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		res := graphViewModel{
			Nodes: []graphNode{},
			Edges: []graphEdge{},
		}

		dataFolder := "./data/1607940066387165600"
		// dataFolder := "./data/1606982502986335200"

		nodesIndex := map[string]graphNode{}
		edgesIndex := map[string]graphEdge{}

		for fi := 0; fi < 50; fi++ {

			edgesFile := fmt.Sprintf("edges_i%v.csv", fi)
			nodesFile := fmt.Sprintf("nodes_i%v.csv", fi)

			f, _ := os.Open(path.Join(dataFolder, nodesFile))
			reader := csv.NewReader(f)
			recs, _ := reader.ReadAll()
			for i, rec := range recs {
				if i == 0 {
					continue
				}

				attr := map[string]string{
					"pc":       rec[2],
					"cn":       rec[3],
					"hi":       rec[4],
					"pi":       rec[5],
					"domscore": rec[6],
				}

				dgr, _ := strconv.Atoi(rec[7])
				init := rec[8] == "true"

				gn := graphNode{
					ID:         rec[0],
					Label:      rec[1],
					Attributes: attr,
					Cluster:    fmt.Sprintf("%v", fi),
					Shared:     false,
					Degree:     dgr,
					IsInit:     init,
				}

				if _, exists := nodesIndex[gn.ID]; exists {
					gn.Shared = true
				}
				nodesIndex[gn.ID] = gn

				// res.Nodes = append(res.Nodes, graphNode{
				// 	ID:         rec[0],
				// 	Label:      rec[1],
				// 	Attributes: attr,
				// 	Cluster:    fmt.Sprintf("%v", i),
				// })
			}

			f, _ = os.Open(path.Join(dataFolder, edgesFile))
			reader = csv.NewReader(f)
			recs, _ = reader.ReadAll()
			for i, rec := range recs {
				if i == 0 {
					continue
				}

				edgesIndex[fmt.Sprintf("%v%v", rec[0], rec[1])] = graphEdge{
					Source: rec[0],
					Target: rec[1],
				}

				// res.Edges = append(res.Edges, graphEdge{
				// 	Source: rec[0],
				// 	Target: rec[1],
				// })
			}

		}

		for _, n := range nodesIndex {
			res.Nodes = append(res.Nodes, n)
		}

		for _, e := range edgesIndex {
			res.Edges = append(res.Edges, e)
		}

		b, _ := json.Marshal(&res)
		w.Write(b)
	}
}

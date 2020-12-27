package graph

import (
	"fmt"
	"testing"

	"github.com/ngeorgiadis/baseline/internal/dataset"
)

func TestKCore(t *testing.T) {

	// g, _ := FromFile("../../data/nodes.csv", "../../data/edges.csv")
	O, _ := FromFile("../../data/nodes.csv", "../../data/edges.csv")

	// top := a.Pop().(*dataset.Author)
	// g[top.ID].Visited = true
	g1 := O.Clone()
	g1 = g1.KCore(22)
	fmt.Println(len(g1))

	for _, n1 := range g1 {
		if n1.Visited {
			fmt.Println(n1.ID)
		}
	}

	max := O.Clone()
	max, k := max.MaxKCore()
	fmt.Println(k)

	fmt.Println(len(max))
	for _, v := range max {
		fmt.Printf("%v: %v\n", v.ID, v.Data.(*dataset.Author).Name)
	}

}

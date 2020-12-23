package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"time"
)

type attr struct {
	ID    string
	Value float64
	Name  string
}

// // AuthorPC ...
// type AuthorPC struct {
// 	ID string
// 	PC int
// }

// // AuthorCN ...
// type AuthorCN struct {
// 	ID string
// 	CN int
// }

// // AuthorHI ...
// type AuthorHI struct {
// 	ID string
// 	HI float64
// }

// // AuthorPI  ...
// type AuthorPI struct {
// 	ID string
// 	PI float64
// }

// AuthDomRow ...
type AuthDomRow struct {
	ID           string
	AuthName     string
	PCSliceIndex int
	CNSliceIndex int
	HISliceIndex int
	PISliceIndex int
	Sum          int
	AprxDomScore int
}

func main() {

	// read input data
	// and construct indexes
	f, _ := os.Open("../../data/nodes_all.csv")
	r := csv.NewReader(f)

	nameIndex := map[string]string{}

	t1 := time.Now()

	fmt.Println("start")

	// nodes (authors)
	pcSlice := []attr{}
	cnSlice := []attr{}
	hiSlice := []attr{}
	piSlice := []attr{}

	fmt.Println("reading...")
	recs, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	for i, row := range recs {
		if i == 0 {
			continue
		}
		id := row[0]
		pc, _ := strconv.Atoi(row[2])
		cn, _ := strconv.Atoi(row[3])
		hi, _ := strconv.ParseFloat(row[4], 64)
		pi, _ := strconv.ParseFloat(row[5], 64)

		pcSlice = append(pcSlice, attr{
			ID:    id,
			Value: float64(pc),
			Name:  "PC",
		})

		cnSlice = append(cnSlice, attr{
			ID:    id,
			Value: float64(cn),
			Name:  "CN",
		})

		hiSlice = append(hiSlice, attr{
			ID:    id,
			Value: hi,
			Name:  "HI",
		})

		piSlice = append(piSlice, attr{
			ID:    id,
			Value: pi,
			Name:  "PI",
		})

		nameIndex[id] = row[1]
	}

	fmt.Println("done")

	fmt.Println("sorting...")

	sort.Slice(pcSlice, func(i, j int) bool {
		return pcSlice[i].Value > pcSlice[j].Value
	})

	sort.Slice(cnSlice, func(i, j int) bool {
		return cnSlice[i].Value > cnSlice[j].Value
	})

	sort.Slice(hiSlice, func(i, j int) bool {
		return hiSlice[i].Value > hiSlice[j].Value
	})

	sort.Slice(piSlice, func(i, j int) bool {
		return piSlice[i].Value > piSlice[j].Value
	})

	fmt.Println("done")

	fmt.Println("creating inv indexes")
	// pcInvIndex := createInvIndex(pcSlice)
	cnInvIndex := createInvIndex(cnSlice)
	hiInvIndex := createInvIndex(hiSlice)
	piInvIndex := createInvIndex(piSlice)
	fmt.Println("done")

	index := map[string]AuthDomRow{}

	n := len(pcSlice)
	fmt.Printf("total elements: %v\n", n)

	fmt.Printf("split time: %v s\n", time.Now().Sub(t1).Seconds())

	fmt.Println("constructing auth dom row index")
	for i := 0; i < n; i++ {

		id := pcSlice[i].ID

		row := AuthDomRow{
			ID:           id,
			PCSliceIndex: i,
			CNSliceIndex: cnInvIndex[id],
			HISliceIndex: hiInvIndex[id],
			PISliceIndex: piInvIndex[id],
			AprxDomScore: -1,
		}

		row.Sum = row.PCSliceIndex + row.CNSliceIndex + row.HISliceIndex + row.PISliceIndex
		// row.AprxDomScore = n - row.Sum

		index[id] = row

		if i%50000 == 0 {
			fmt.Printf("split 50k time: %v s\n", time.Now().Sub(t1).Seconds())
		}
	}
	fmt.Println("done")

	fmt.Printf("split time: %v s\n", time.Now().Sub(t1).Seconds())
	fmt.Println("approximating domscrore")

	i := 0

	for k, v := range index {

		if v.PCSliceIndex > 5000 &&
			v.CNSliceIndex > 5000 &&
			v.HISliceIndex > 5000 &&
			v.PISliceIndex > 5000 {
			continue
		}

		// a := pcSlice[0:v.PCSliceIndex]
		// b := cnSlice[0:v.CNSliceIndex]
		// c := hiSlice[0:v.HISliceIndex]
		// d := piSlice[0:v.PISliceIndex]

		x := countDistinct(v, pcSlice, cnSlice, hiSlice, piSlice)

		v.AprxDomScore = n - x
		index[k] = v

		i++

		if i%1000 == 0 {
			fmt.Printf("split 1k time: %v s\n", time.Now().Sub(t1).Seconds())
		}

	}

	fmt.Println("done")

	fmt.Println("writing to file")

	of, _ := os.Create(fmt.Sprintf("../../data/apprxDom.csv"))
	writer := csv.NewWriter(of)

	writer.Write([]string{
		"ID",
		"Name",
		"PC index",
		"CN index",
		"HI index",
		"PI index",
		"Approx. Score",
	})

	sortedRes := []AuthDomRow{}
	for _, v := range index {
		sortedRes = append(sortedRes, v)
	}

	sort.Slice(sortedRes, func(i, j int) bool {
		return sortedRes[i].AprxDomScore > sortedRes[j].AprxDomScore
	})

	for _, v := range sortedRes {
		row := []string{}

		row = append(row, v.ID)
		row = append(row, nameIndex[v.ID])
		row = append(row, fmt.Sprintf("%v", v.PCSliceIndex))
		row = append(row, fmt.Sprintf("%v", v.CNSliceIndex))
		row = append(row, fmt.Sprintf("%v", v.HISliceIndex))
		row = append(row, fmt.Sprintf("%v", v.PISliceIndex))
		row = append(row, fmt.Sprintf("%v", v.AprxDomScore))

		writer.Write(row)
	}

	// for i, row := range records {
	// 	if i == 0 {
	// 		row = append(row, "domscore")
	// 		row = append(row, "degree")
	// 		row = append(row, "is_initial")
	// 		writer.Write(row)
	// 		continue
	// 	}

	// 	if n, ok := ixG[row[0]]; ok {
	// 		row = append(row, fmt.Sprintf("%v", n.Data.GetDomScore()))
	// 		row = append(row, fmt.Sprintf("%v", len(n.Neighbours)))
	// 		row = append(row, fmt.Sprintf("%v", n.IsInitialCommunityNode))
	// 		writer.Write(row)
	// 	}
	// }

	writer.Flush()
	of.Close()

	fmt.Println("done")
	fmt.Println("finish")

	fmt.Printf("total time: %v s\n", time.Now().Sub(t1).Seconds())

}

func countDistinct(v AuthDomRow, arrays ...[]attr) int {

	index := map[string]string{}
	// for _, a := range arrays {
	// 	for _, b := range a {
	// 		index[b.ID] = ""
	// 	}
	// }
	for _, b := range arrays[0][0:v.PCSliceIndex] {
		index[b.ID] = ""
	}

	for _, b := range arrays[1][0:v.CNSliceIndex] {
		index[b.ID] = ""
	}

	for _, b := range arrays[2][0:v.HISliceIndex] {
		index[b.ID] = ""
	}

	for _, b := range arrays[3][0:v.PISliceIndex] {
		index[b.ID] = ""
	}

	return len(index)

}

func findIndex(s []attr, id string) int {
	for i, r := range s {
		if r.ID == id {
			return i
		}
	}

	return -1
}

func createInvIndex(s []attr) map[string]int {
	index := map[string]int{}
	for i, v := range s {
		index[v.ID] = i
	}
	return index
}

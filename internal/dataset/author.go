package dataset

import "container/heap"

// Author ...
type Author struct {
	ID                string
	Name              string
	PublicationNumber int
	CitationNumber    int
	HIndex            float64
	PIndex            float64
	DomScore          int
	Visited           bool

	index int
}

//implementation of the DominanceReaderWriter interface

// GetDomScore ...
func (a *Author) GetDomScore() int {
	return a.DomScore
}

// SetDomScore ...
func (a *Author) SetDomScore(v int) {
	a.DomScore = v
}

// AuthorPriorityQueue ...
type AuthorPriorityQueue []*Author

func (apq AuthorPriorityQueue) Len() int { return len(apq) }

func (apq AuthorPriorityQueue) Less(i, j int) bool {
	return apq[i].DomScore > apq[j].DomScore
}

func (apq AuthorPriorityQueue) Swap(i, j int) {
	apq[i], apq[j] = apq[j], apq[i]
	apq[i].index = i
	apq[j].index = j
}

// Push ...
func (apq *AuthorPriorityQueue) Push(x interface{}) {
	n := len(*apq)
	item := x.(*Author)
	item.index = n
	*apq = append(*apq, item)
}

// Pop ...
func (apq *AuthorPriorityQueue) Pop() interface{} {
	old := *apq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*apq = old[0 : n-1]
	return item
}

// Update ...
func (apq *AuthorPriorityQueue) Update(item *Author, score int) {
	item.DomScore = score
	heap.Fix(apq, item.index)
}

// Dominates ...
func (a *Author) Dominates(b *Author) bool {
	domRule1 := a.CitationNumber >= b.CitationNumber
	domRule2 := a.PublicationNumber >= b.PublicationNumber
	domRule3 := a.HIndex >= b.HIndex
	domRule4 := a.PIndex >= b.PIndex

	if a.CitationNumber == b.CitationNumber &&
		a.PublicationNumber == b.PublicationNumber &&
		a.HIndex == b.HIndex &&
		a.PIndex == b.PIndex {
		return false
	}

	return domRule1 && domRule2 && domRule3 && domRule4
}

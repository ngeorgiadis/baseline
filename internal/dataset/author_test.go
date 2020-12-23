package dataset

import (
	"container/heap"
	"fmt"
	"testing"
)

func TestAuthorPQ(t *testing.T) {

	// items :=

	apq := AuthorPriorityQueue{
		&Author{
			Name:     "nick",
			DomScore: 10,
		},
		&Author{
			Name:     "George",
			DomScore: 15,
		},
		&Author{
			Name:     "Kathrin",
			DomScore: 21,
		},
	}

	heap.Init(&apq)
	a := &Author{
		Name:     "nick2",
		DomScore: 5,
	}
	heap.Push(&apq, a)
	apq.Update(a, 33)

	for apq.Len() > 0 {
		item := heap.Pop(&apq).(*Author)
		fmt.Printf("%.2d:%s ", item.DomScore, item.Name)
	}

}

func TestAuthorInterface(t *testing.T) {
	a := &Author{}

	a.GetDomScore()
}

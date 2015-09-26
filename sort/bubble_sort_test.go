package sort

import (
	"fmt"
	"testing"
)

type Sortable []int64

func (s Sortable) Len() int {
	return len(s)
}

func (s Sortable) Less(indF, indL int) bool {
	return s[indF] < s[indL]
}

func (s Sortable) Equal(indF, indL int) bool {
	return s[indF] == s[indL]
}

func (s Sortable) Swap(indF, indL int) {
	s[indF], s[indL] = s[indL], s[indF]
}

func sortableInts() Sortable {
	var sortableInts Sortable
	sortableInts = []int64{74, 59, 238, -784, 9845, 959, 905, 0, 0, 42, 7586, -5467984, 7586}
	return sortableInts
}

func isSortedDsc(arr Sortable) bool {
	for ind := 1; ind < arr.Len(); ind++ {
		if arr.Less(ind, ind-1) == false && arr.Equal(ind, ind-1) == false {
			return false
		}
	}
	return true
}

func isSortedAsc(arr Sortable) bool {
	for ind := 1; ind < arr.Len(); ind++ {
		if arr.Less(ind, ind-1) == true && arr.Equal(ind, ind-1) == false {
			return false
		}
	}
	return true
}

func TestBubbleSortAscendingInts(t *testing.T) {
	data := sortableInts()
	BubbleAsc(data)
	if isSortedAsc(data) != true {
		t.Error(fmt.Sprintln("Tried to sort", sortableInts(), "in asc, but got", data))
	}
}

func TestBubbleSortDescendingInts(t *testing.T) {
	data := sortableInts()
	BubbleDsc(data)
	if isSortedDsc(data) != true {
		t.Error(fmt.Sprintln("Tried to sort", sortableInts(), "in dsc, but got", data))
	}
}

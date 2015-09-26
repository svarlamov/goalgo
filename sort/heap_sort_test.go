package sort

import (
	"fmt"
	"testing"
)

func TestHeapSortAscendingInts(t *testing.T) {
	data := sortableInts()
	HeapSortAsc(data)
	if isSortedAsc(data) != true {
		t.Error(fmt.Sprintln("Tried to sort", sortableInts(), "in asc, but got", data))
	}
}

func TestHeapSortDescendingInts(t *testing.T) {
	data := sortableInts()
	HeapSortDsc(data)
	if isSortedDsc(data) != true {
		t.Error(fmt.Sprintln("Tried to sort", sortableInts(), "in dsc, but got", data))
	}
}

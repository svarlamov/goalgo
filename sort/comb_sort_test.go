package sort

import (
	"fmt"
	"testing"
)

func TestCombSortAscendingInts(t *testing.T) {
	data := sortableInts()
	CombSortAsc(data)
	if isSortedAsc(data) != true {
		t.Error(fmt.Sprintln("Tried to sort", sortableInts(), "in asc, but got", data))
	}
}

func TestCombSortDescendingInts(t *testing.T) {
	data := sortableInts()
	CombSortDsc(data)
	if isSortedDsc(data) != true {
		t.Error(fmt.Sprintln("Tried to sort", sortableInts(), "in dsc, but got", data))
	}
}

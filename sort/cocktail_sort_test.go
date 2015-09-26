package sort

import (
	"fmt"
	"testing"
)

func TestCocktailSortAscendingInts(t *testing.T) {
	data := sortableInts()
	CocktailAsc(data)
	if isSortedAsc(data) != true {
		t.Error(fmt.Sprintln("Tried to sort", sortableInts(), "in asc, but got", data))
	}
}

func TestCocktailSortDescendingInts(t *testing.T) {
	data := sortableInts()
	CocktailDsc(data)
	if isSortedDsc(data) != true {
		t.Error(fmt.Sprintln("Tried to sort", sortableInts(), "in dsc, but got", data))
	}
}

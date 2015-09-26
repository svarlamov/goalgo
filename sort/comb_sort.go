package sort

func CombSortDsc(arr Comparable) {
	combSort(arr, true)
}

func CombSortAsc(arr Comparable) {
	combSort(arr, false)
}

func combSort(arr Comparable, dsc bool) {
	gap := arr.Len()
	for {
		if gap > 1 {
			gap = gap * 100 / 124
		}
		for i := 0; ; {
			if arr.Less(i+gap, i) == !dsc {
				arr.Swap(i+gap, i)
			}
			i++
			if i+gap >= arr.Len() {
				break
			}
		}
		if gap == 1 {
			break
		}
	}
}

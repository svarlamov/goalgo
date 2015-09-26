package sort

func HeapSortAsc(arr Comparable) {
	heapSort(arr, false)
}

func HeapSortDsc(arr Comparable) {
	heapSort(arr, true)
}

func sift(arr Comparable, i int, arrLen int, dsc bool) {
	done := false
	maxChild := 0

	for (i*2+1 < arrLen) && (!done) {
		if i*2+1 == arrLen-1 {
			maxChild = i*2 + 1
		} else if arr.Less(i*2+2, i*2+1) != dsc {
			maxChild = i*2 + 1
		} else {
			maxChild = i*2 + 2
		}

		if arr.Less(i, maxChild) != dsc {
			arr.Swap(i, maxChild)
			i = maxChild
		} else {
			done = true
		}
	}
}

func heapSort(arr Comparable, dsc bool) {
	i := 0

	for i = arr.Len()/2 - 1; i >= 0; i-- {
		sift(arr, i, arr.Len(), dsc)
	}

	for i = arr.Len() - 1; i >= 1; i-- {
		arr.Swap(0, i)
		sift(arr, 0, i, dsc)
	}
}

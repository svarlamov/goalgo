package sort

func BubbleAsc(arr Comparable) {
	bubble(arr, false)
}

func BubbleDsc(arr Comparable) {
	bubble(arr, true)
}

func bubble(arr Comparable, dsc bool) {
	for i := 0; i < arr.Len(); i++ {
		for j := 0; j < arr.Len()-1; j++ {
			if arr.Less(j, j+1) == dsc {
				arr.Swap(j, j+1)
			} else if arr.Less(j, j+1) == dsc {
				arr.Swap(j, j+1)
			}
		}
	}
}

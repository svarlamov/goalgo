package sort

func CocktailDsc(arr Comparable) {
	cocktail(arr, true)
}

func CocktailAsc(arr Comparable) {
	cocktail(arr, false)
}

func cocktail(arr Comparable, dsc bool) {
	for i := 0; i < arr.Len()/2; i++ {
		left := 0
		right := arr.Len() - 1
		for left <= right {
			if arr.Less(left+1, left) == !dsc {
				arr.Swap(left, left+1)
			}
			left++
			if arr.Less(right, right-1) == !dsc {
				arr.Swap(right, right-1)
			}
			right--
		}
	}
}

package sort

type Comparable interface {
	Len() int
	Less(int, int) bool
	Swap(int, int)
}

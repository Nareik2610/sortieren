package bubblesort

// BubbleSort sortiert the gegebene Liste mit dem Bubble-Sort-Algorithmus.
func BubbleSort(list []int) {
	// TODO
	for i := len(list) - 1; i > 0; i-- {
		BubbleUp(list[:i+1])
	}
}

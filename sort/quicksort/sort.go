package quicksort

import (
	"fmt"
	"github.com/varminas/patterns-for-coding-questions/utils"
)

// 5, 1, 4, 2, 8
func Quicksort(arr []int) []int {
	quicksort(arr, 0, len(arr)-1)
	return arr
}

func quicksort(arr []int, low int, high int) {
	//fmt.Printf("%v\n", arr)
	if low < high {
		pi := partition(arr, low, high)
		fmt.Printf("arr[pi]=%d, pi=%d\n", arr[pi], pi)

		quicksort(arr, low, pi-1)
		quicksort(arr, pi+1, high)
	}
}

func partition(arr []int, low int, high int) int {
	pivot := arr[high]

	i := low - 1
	for j := low; j <= high-1; j++ {
		if arr[j] < pivot {
			i++
			utils.Swap2(arr, i, j)
		}
	}
	utils.Swap2(arr, i+1, high)
	return i + 1
}

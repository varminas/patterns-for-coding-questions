package bubble

import "fmt"

// 5, 1, 4, 2, 8
func Sort(arr []int) []int {
	count := 0
	swapped := true
	n := len(arr)
	for swapped == true {
		swapped = false
		for i := 1; i < n-1; i++ {
			count++
			if arr[i-1] > arr[i] {
				swap(arr, i)
				swapped = true
			}
		}
		n = n - 1
	}
	fmt.Printf("Iterations: %d\n", count)
	return arr
}

func swap(arr []int, i int) {
	tmp := arr[i-1]
	arr[i-1] = arr[i]
	arr[i] = tmp
}

package mergesort

import "fmt"

var count = 1

func Mergesort(arr []int) []int {
	sort(arr, len(arr))
	return arr
}

func sort(arr []int, n int) {
	fmt.Printf("SORT. count=%d, arr=%v\n", count, arr)
	count++
	if n < 2 {
		return
	}

	// +1 splits array of odd (e.g. [1,2,3,4,5,6,7] to [1,2,3,4][5,6,7]
	mid := (n + 1) / 2
	l := make([]int, mid)
	r := make([]int, n-mid)
	for i := 0; i < mid; i++ {
		l[i] = arr[i]
	}
	for i := mid; i < n; i++ {
		r[i-mid] = arr[i]
	}

	sort(l, mid)
	sort(r, n-mid)
	merge(arr, l, r, mid, n-mid)
}

func merge(arr []int, l []int, r []int, left int, right int) {
	fmt.Printf("MERGE. count=%d, arr=%v\n", count, arr)
	count++
	i := 0
	j := 0
	k := 0

	for i < left && j < right {
		if l[i] <= r[j] {
			arr[k] = l[i]
			i++
			k++
		} else {
			arr[k] = r[j]
			j++
			k++
		}
	}

	for i < left {
		arr[k] = l[i]
		i++
		k++
	}

	for j < right {
		arr[k] = r[j]
		j++
		k++
	}
}

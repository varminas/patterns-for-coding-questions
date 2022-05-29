package utils

import "fmt"

func PrintErrorString(funcName string, input interface{}, got interface{}, want interface{}) string {
	return fmt.Sprintf("%v(%#v) \n\ngot= %#v \n\nwant=%#v", funcName, input, got, want)
}

func Swap(arr []int, i int) {
	tmp := arr[i-1]
	arr[i-1] = arr[i]
	arr[i] = tmp
}

func Swap2(arr []int, i int, j int) {
	tmp := arr[j]
	arr[j] = arr[i]
	arr[i] = tmp
}

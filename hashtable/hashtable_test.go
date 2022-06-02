package hashtable

import (
	"fmt"
	"testing"
)

func TestHashtable(t *testing.T) {
	myMap := &Map{}
	myMap.Init()
	fmt.Printf("IsEmpty: %v\n", myMap.IsEmpty())
	myMap.Add("golang", "GoLang language")
	myMap.Add("cpp", "C++")
	fmt.Printf("IsEmpty: %v\n", myMap.IsEmpty())

	fmt.Printf("Get by key \"golang\". Value=%v\n", myMap.Get("golang"))
	fmt.Printf("Get by key \"cpp\". Value=%v\n", myMap.Get("cpp"))
	fmt.Printf("%v", myMap)

	//input := []int{10, 80, 30, 90, 40, 50, 70}
	//got := Main()
	//fmt.Printf("Quicksort result: %v", input)
	//want := []int{10, 30, 40, 50, 70, 80, 90}
	//
	//if !reflect.DeepEqual(got, want) {
	//	t.Errorf(utils.PrintErrorString("HashTable", input, got, want))
	//}
}

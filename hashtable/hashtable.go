package hashtable

// Separate Chaining Technique
type HashNode struct {
	Key      string
	Value    interface{}
	HashCode int
	Next     *HashNode
}

type Map struct {
	bucketArray []*HashNode
	numBuckets  int
	size        int
}

func (myMap *Map) Init() {
	myMap.bucketArray = make([]*HashNode, 10)
	myMap.numBuckets = 10
	myMap.size = 0

	for i := 0; i < myMap.numBuckets; i++ {
		//myMap.bucketArray[i] = nil
	}
}

func (myMap *Map) GetSize() int {
	return myMap.size
}

func (myMap *Map) IsEmpty() bool {
	return myMap.size == 0
}

func (myMap *Map) hashCode(key string) int {
	var sum int
	for _, v := range key {
		sum += int(v)
	}
	return sum
}

func (myMap *Map) getBucketIndex(key string) int {
	hashCode := myMap.hashCode(key)
	index := hashCode % myMap.numBuckets
	if index < 0 {
		index = index * -1
	}
	return index
}

func (myMap *Map) Add(key string, value interface{}) {
	bucketIndex := myMap.getBucketIndex(key)
	hashCode := myMap.hashCode(key)
	head := myMap.bucketArray[bucketIndex]

	for head != nil {
		if head.Key == key && hashCode == head.HashCode {
			return
		}
		head = head.Next
	}

	myMap.size++
	head = myMap.bucketArray[bucketIndex]
	newNode := &HashNode{Key: key, Value: value, HashCode: hashCode}
	newNode.Next = head
	myMap.bucketArray[bucketIndex] = newNode

	factor := float64(1.0*myMap.size) / float64(myMap.numBuckets)
	if factor >= 0.7 {
		temp := myMap.bucketArray
		myMap.bucketArray = make([]*HashNode, 1)
		myMap.numBuckets = 2 * myMap.numBuckets
		myMap.size = 0
		for i := 0; i < myMap.numBuckets; i++ {
			myMap.bucketArray = append(myMap.bucketArray, nil)
		}

		for _, headNode := range temp {
			for headNode != nil {
				myMap.Add(headNode.Key, headNode.Value)
				headNode = headNode.Next
			}
		}
	}
}

func (myMap *Map) Get(key string) interface{} {
	bucketIndex := myMap.getBucketIndex(key)
	hashCode := myMap.hashCode(key)

	head := myMap.bucketArray[bucketIndex]

	for head != nil {
		if head.Key == key && head.HashCode == hashCode {
			return head.Value
		}
		head = head.Next
	}
	return nil
}

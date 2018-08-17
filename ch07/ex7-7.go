// heap.Interface 타입

type Interface interface {
	sort.Interface
	Push(x interface{})
	Pop() interface{}
}

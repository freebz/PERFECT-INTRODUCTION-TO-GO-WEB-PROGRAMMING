// sort.Interface 타입

type Interface interface {
	Len() int
	Less(i, j int) bool
	Swap(i, j int)
}

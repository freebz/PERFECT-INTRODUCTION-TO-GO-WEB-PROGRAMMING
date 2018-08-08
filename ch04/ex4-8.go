// 리시버 변수 생략

type rect struct {
	width  float64
	height float64
}

func (rect) new() rect {
	return rect{}
}

func main() {
	r := rect{}.new()
	fmt.Println(r)
}

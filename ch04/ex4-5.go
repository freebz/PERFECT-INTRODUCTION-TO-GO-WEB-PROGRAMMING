// 인터페이스

type shaper interface {
	area() float64
}

func describe(s shaper) {
	fmt.Println("area :", s.area())
}

func main() {
	r := rect{3, 4}
	describe(r)
}

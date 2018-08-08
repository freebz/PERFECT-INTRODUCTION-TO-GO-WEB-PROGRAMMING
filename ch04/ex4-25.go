// 빈 인터페이스

func main() {
	r := rect{3, 4}
	c := circle{2.5}
	display(r)              // width: 3.000000, height: 4.000000
	display(c)              // radius: 2.500000
	display(2.5)            // 2.5
	display("rect struct")  // rect struct
}

func display(s interface{}) {
	fmt.Println(s)
}

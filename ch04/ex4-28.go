// 기본 라이브러리의 인터페이스 활용 - fmt.Stringer

func Println(a ...interface{}) (n int, err error) {
	return Fprintln(os.Stdout, a...)
}

func Fprintf(w io.Writer, a ...interface{}) (n int, err error) {
	...
	// fmt.Stringer 인터페이스 타입일 때 String() 메서드의 결괏값을출력
	...
	return
}


type Stringer interface {
	String() string
}



func (t Item) String() string {
	return fmt.Sprintf("[%s] %.0f", t.name, t.Cost())
}

func (t DiscountItem) String() string {
	return fmt.Sprintf("%s => %.0f((%.0f%s DC)", t.Item.String(), t.Cost(), t.discountRate, "%")
}

func (t Rental) String() string {
	return fmt.Sprintf("[%s] %.0f", t.name, t.Cost())
}

func (ts Items) String() string {
	var s []string
	for _, t := range ts {
		s = append(s, fmt.Sprint(t))
	}
	return fmt.Sprintf("%d items. total: %.0f\n\t- %s",
		len(ts), ts.Cost(), strings.Join(s, "\n\t- "))
}



func main() {
	shirt := Item("Men's Slim-Fit Shirt", 25000, 3}
	video := Rental{"Interstellar", 1000, 3, Days}
	eventShoes := DiscountItem{
		Item{"Women's Walking Shoes", 50000, 3},
		10.00,
	}
	items := Items{shirt, video, eventShoes}

	fmt.Println(shirt)
	fmt.Println(eventSheets)
	fmt.Println(items)
}

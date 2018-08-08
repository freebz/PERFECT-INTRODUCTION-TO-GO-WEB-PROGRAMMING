// 인터페이스 임베딩

type Itemer interface {
	Coster
	fmt.Stringer
}


type Order struct {
	Itemer
	taxRate float64
}

func (o Order) Cost() float64 {
	return o.Itemer.Cost() * (1.0 + o.taxRate/100)
}

func (o Order) String() string {
	return fmt.Sprintf("Total price: %.0f(tax rate: %.2f)\n\tOrder details: %s",
		o.Cost(), o.taxRate, o.Itemer.String())
}


func main() {
	shirt := Item{"Men's Slim-Fit Shirt", 25000, 3}
	video := Rental{"Interstellar", 1000, 3, Days}
	eventShoes := DiscountItem{
		Item{"Women's Walking Shoes", 50000, 3},
		10.00,
	}

	order1 := Order{Items{shirt, eventShoes}, 10.00}
	order2 := Order{video, 5.00}

	fmt.Println(order1)
	fmt.Println(order2)
}

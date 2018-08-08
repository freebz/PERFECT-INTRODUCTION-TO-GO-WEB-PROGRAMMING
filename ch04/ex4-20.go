// 메서드 오버라이딩

func (t DiscountItem) Cost() float64 {
	return t.Item.Cost() * (1.0 - t.discountRate/100)
}

func main() {
	shoes := Item{"Women's Walking Shoes", 30000, 2}
	eventShoes := DiscountItem{
		Item{"Sports Shoes", 50000, 3},
		10.00,
	}

	fmt.Println(shoes.Cost())           // 60000
	fmt.Println(eventShoes.Cost())      // 135000
	fmt.Println(eventShoes.Item.Cost()) // 150000
}

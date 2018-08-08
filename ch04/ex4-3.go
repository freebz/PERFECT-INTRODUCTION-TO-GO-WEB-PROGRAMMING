// 함수 서명을 사용자 정의 타입으로 사용

type quantity int
type costCalcultor func(quantity, float64) float64

func describe(q quentity, price float64, c costCalculator) {
	fmt.Printf("quantity: %d, price: %0.5f, cost: %0.0f\n",
		q, price, c(q, price))
}

func main() {
	var offBy10Percent, offBy100Won costCalculator

	offBy10Percent = func(q quantity, price float64) float64 {
		return float64(q) * price * 0.9
	}
	offBy1000Won = func(q quantity, price float64) float64 {
		return float64(q)*price - 1000
	}

	describe(3, 10000, offBy10Percent)
	describe(3, 10000, offBy1000Won)
}

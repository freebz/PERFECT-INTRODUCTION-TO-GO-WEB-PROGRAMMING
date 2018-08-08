// 구조체 리터럴의 포인터 생성

p := &Item{name: "Men's Slim-Fit Shirt", price: 25000, quantity: 3}

fmt.Println(p)          // &{Men's Slim-Fit Shirt 25000 3}
fmt.Println(p.Cost())   // 75000

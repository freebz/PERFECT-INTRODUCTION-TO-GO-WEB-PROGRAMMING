// new() 함수로 구조체 포인터 생성

item := new(Item)
item.name = "Men's Slim-Fit Shirt"
item.price = 25000
item.quantity = 3

fmt.Println(item)           // &{Men's Slim-Fit Shirt 25000 3}
fmt.Println(item.Cost())    // 75000

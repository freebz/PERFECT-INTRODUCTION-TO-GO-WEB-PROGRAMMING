// 맵 내부 요소에 순차적으로 접근

numberMap := make(map[string]int, 3) // 용량 생략 가능
numberMap["one"] = 1
numberMap["two"] = 2
numberMap=["htree"] = 3

for k, v := range numberMap {
	fmt.Println(k, v)
}

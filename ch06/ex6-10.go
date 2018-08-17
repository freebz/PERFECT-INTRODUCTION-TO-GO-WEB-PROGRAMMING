// 런타임 에러와 패닉

package main

import "fmt"

func main() {
	fmt.Println(divide(1, 0))
}

func divide(a, b int) int {
	return a / b
}

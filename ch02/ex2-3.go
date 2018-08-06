// for 반복문 2

package main

import "fmt"

func main() {
	sum, i := 0, 0
	// for 문에 조건식만 사용
	for i < 10 {
		sum += i
		i++
	}
	fmt.Println(sum)
}

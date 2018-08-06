// for 반복문 3

package main

import "fmt"

func main() {
	sum, i := 0, 0
	// for 문에 조건식 생략
	for {
		if i >= 10 {
			break
		}
		sum += i
		i++
	}
	fmt.Println(sum)
}

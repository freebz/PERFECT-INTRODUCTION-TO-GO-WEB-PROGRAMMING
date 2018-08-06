// 증감 연산은 반환 값이 없음

package main

import "fmt"

func main() {
	sum, i := 0, 0
	for i < 10 {
		sum += i++ // 증감 연산은 반환 값이 없음. 컴파일 에러 발생
	}
	fmt.Println(sum)
}

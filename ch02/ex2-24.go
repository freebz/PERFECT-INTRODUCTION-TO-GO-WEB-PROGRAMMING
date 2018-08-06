// strconv.Atoi 함수의 정상 수행 여부에 따라 각각 다른 로직 처리

package main

import (
	"fmt"
	"strconv"
)

func displayInt(s string) {
	if v, err := strconv.Atoi(s); err != nil {
		fmt.Printf("%s는 정수가 아닙니다.\n", s)
	} else {
		fmt.Printf("정수 값은 %d입니다.\n", v)
	}
}

func main() {
	displayInt("two")
	displayInt("2")
}

package main

import (
	"fmt"
	mylib "go-book-sample/ch2/lib"
)

func main() {
	fmt.Println(mylib.IsDigit('1')) // lib 패키지의 IsDigit 함수 사용
	fmt.Println(mylib.IsDigit('a')) // lib 패키지의 IsDigit 함수 사용
}

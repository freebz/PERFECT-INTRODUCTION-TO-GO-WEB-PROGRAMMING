// $GOPATH/src/go-book-sample/ch2/pkg/main.go

package main

import (
	"fmt"
	"go-book-sample/ch2/lib"
)

func main() {
	fmt.Println(lib.IsDigit('1')) // lib 패키지의 IsDigit 함수 사용
	fmt.Println(lib.IsDigit('a')) // lib 패키지의 IsDigit 함수 사용
}

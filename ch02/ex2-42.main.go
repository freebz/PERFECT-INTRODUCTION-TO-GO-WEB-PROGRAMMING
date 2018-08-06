// $GOPATH/src/go-book-sample/ch2/pkg2/main.go

package main

import (
	"fmt"
	"go-book-sample/ch2/lib"
)

var v rune

func init() {
	v = '1'
}

func main() {
	fmt.Println(lib.IsDigit(v)) // lib 패키지의 IsDigit 함수 사용
}

func isDigit(c int32) bool {
	return '0' <= c && c <= '9'
}

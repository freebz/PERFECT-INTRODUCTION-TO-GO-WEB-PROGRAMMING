// $GOPATH/src/go-book-sample/ch2/kg/main.go

package main

import (
	"fmt"
	"go-book-sample/ch2/lib"
)

func main() {
	fmt.Println(lib.IsDigit('1'))
	fmt.Println(lib.IsDigit('a'))
	fmt.Println(lib.isSpace('\t'))
}

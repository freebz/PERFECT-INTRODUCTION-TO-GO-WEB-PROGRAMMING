// 참조에 의한 호출

package main

import "fmt"

func inc(i *int) {
	*i = *i + 1
}

func main() {
	i := 10
	inc(&i)
	fmt.Println(i)
}

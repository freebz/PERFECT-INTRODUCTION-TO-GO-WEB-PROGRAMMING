// defer를 사용하여 트레이스 로그 출력 2

package main

import "fmt"

func enter(s string) string {
	fmt.Println("entering:", s)
	return s
}
func leave(s string) {
	fmt.Println("leaving:", s)
}
func a() {
	defer leave(enter("a"))
	fmt.Println("in a")
}
func b() {
	defer leave(enter("b"))
	fmt.Println("in b")
	a()
}
func main() {
	b()
}

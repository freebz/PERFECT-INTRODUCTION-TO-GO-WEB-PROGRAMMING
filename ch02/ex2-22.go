// 반환 값이 두 개 이상일 때

package main

import "fmt"

func multiply(w, h int) (int, int) {
	return w * 2, h * 2
}

func main() {
	w, h := multiply(3, 4)
	fmt.Println(w, h)
}

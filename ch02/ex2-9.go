// 스타일이 맞지 않는 코드

package main
import "fmt"
func main() {
	type Rect struct{
	width int //width
	height int //heigh
	}

	r := Rect{1, 2}; fmt.Println(r.width * 2+r.height * 2)
}

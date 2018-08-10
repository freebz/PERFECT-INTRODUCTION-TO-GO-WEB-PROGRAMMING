// 버퍼트 채널

package main

import "fmt"

func main() {
	c := make(chan int, 2)
	c <- 1
	c <- 2
	c <- 3
	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c)
}

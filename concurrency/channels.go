package main

import "fmt"

func main() {
	c := make(chan int, 5)

	c <- 1
	c <- 2
	c <- 3
	c <- 4

	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c)

}

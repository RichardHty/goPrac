package main

import "fmt"

func worker(c chan int) {
	for {
		n := <-c
		fmt.Println(n)
	}

}
func chanDemo() {
	c := make(chan int)
	go worker(c)
	c <- 1
	c <- 2

}

func main() {
	chanDemo()
}

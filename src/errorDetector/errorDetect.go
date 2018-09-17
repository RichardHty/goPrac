package main

import (
	"bufio"
	"fmt"
	"os"
)

func tryDefer() {
	// defer stack
	defer fmt.Println(1)
	defer fmt.Println(2)
	fmt.Println(3)
	return
}
func writeFile(fileName string) {
	file, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	defer writer.Flush()

	for i := 0; i < 10; i++ {
		fmt.Fprintln(writer, i)
	}
}
func main() {
	tryDefer()
	writeFile("num.txt")
}

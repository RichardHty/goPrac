package main

import (
	"fmt"
	"math"
	"io/ioutil"
	"reflect"
	"runtime"
	"strconv"
)

const Filename = "abc.txt"

func variableZeroValue() {
	var a, b  = 3, 4
	var s string
	fmt.Printf("%d, %q, %d", a, s, a+b)
}

// conts & enum
func consts() {
	const (
		Filename = "abc.txt"
		a, b     = 3, 4
	)
	// iota: increment
	const (
		kb = 1 << (2* iota)
		mb
		gb
		tb
		pb
	)

	fmt.Println(kb,mb,gb,tb,pb)
	var c = int( math.Sqrt(a*a + b*b) )
	fmt.Println(Filename, c)
}

// switch prac
func grade(score int) string {
	g := ""
	switch {
	case score < 0 || score > 100:
		panic(fmt.Sprintf(
			"Wrong score: %d", score))
	case score < 60:
		g = "F"
	case score < 80:
		g = "C"
	case score < 90:
		g = "B"
	case score <= 100:
		g = "A"

	}
	return g
}

//for loop prac
func convertToBin(value int) string {
	result := ""
	if value <= 0 {
		return "0"
	}
	for ; value > 0 ; value /= 2{
		lsb := value % 2
		result = strconv.Itoa(lsb) + result
	}
	return result
}

// pass arg
func swap(a, b *int) {
	*b, *a = *a, *b
	return
}

// function, error handler
func eval(a, b int, op rune) (int, error) {

	switch op {
	case '+':
		return a + b, nil
	case '-':
		return a - b, nil
	default:
		return 0, fmt.Errorf(
			"unsupported operation: %s" , string(op))
	}
}
// func oriented
func apply(op func(int, int) int, a, b int) int {
	p := reflect.ValueOf(op).Pointer()
	opName := runtime.FuncForPC(p).Name()
	fmt.Printf("Calling function %s with args: " +
		"(%d, %d)", opName, a, b)
	return op(a, b)
}
func pow(a, b int) int{
	return int(math.Pow(float64(a),float64(b)))
}

// number of args is variable
func sumArgs(values ... int) int {
	sum := 0
	for i := range values {
		sum += values[i]
	}
	return sum
}
func basic() {
	fmt.Println("Good")
	fmt.Println(Filename)
	consts()
	fmt.Println(Filename)
	variableZeroValue()

	// read file
	if  contents,err := ioutil.ReadFile(Filename); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s\n", contents)
	}

	fmt.Println(convertToBin(1))
	fmt.Println(grade(0))

	// swap
	a, b := 2, 3
	swap(&a,&b)
	fmt.Println(a, b)

	// eval
	fmt.Println(eval(1,2,'n'))

	// apply
	fmt.Println(apply(pow,10,3))

	// sum args
	fmt.Println(sumArgs(1,2,3,4))

}

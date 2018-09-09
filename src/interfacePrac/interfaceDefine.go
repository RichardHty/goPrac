package main

import (
	"fmt"
	"interfacePrac/mock"
	real2 "interfacePrac/real"
	"time"
)

type Retriever interface {
	Get(url string) string
}
func download(r Retriever) string {
	return r.Get("http://www.imooc.com")
}
func main () {
	var r Retriever
	r = mock.Retriever{"This is a fake content"}
	inspect(r)
	fmt.Println()

	// type assertion
	if retriver,ok := r.(*real2.Retriever); ok{
		fmt.Println(retriver.UserAgent)
	} else {
		fmt.Println("not a *real retriever!")
	}


	r = &real2.Retriever{
		UserAgent:"Mozilla/5.0",
		TimeOut:time.Minute,
	}
	//fmt.Printf("%T %v",r, r)

	inspect(r)

	retriver := r.(*real2.Retriever)
	fmt.Println(retriver.TimeOut)

}

func inspect(r Retriever) {
	switch v := r.(type) {
	case mock.Retriever:
		fmt.Println("Contents: ", v.Contents)
	case *real2.Retriever:
		fmt.Println("UserAgent: ", v.UserAgent)
	}
}

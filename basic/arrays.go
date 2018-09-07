package main

import "fmt"

// range key word
func sum(numbers [3]int) int{
	sum := 0
	for _,v := range numbers {
		sum += v
		fmt.Println(v)
	}
	return sum
}

// slice: changes on slice will be reflected on original array
func updateSlice(s []int) {
	s[0] = 100
}
func main() {
	arr := [...]int{1,2,3}
	fmt.Println(sum(arr))

	arr2 := [...]int{0,1,2,3,4,5,6,7}

	// slice is [...)
	s1 := arr2[:6]
	fmt.Println("This is s1:=arr2[:6] ")
	fmt.Println(s1)

	s2 := arr2[2:]
	fmt.Println("This is s2:=arr2[2:] ")
	fmt.Println(s2)

	s3 := arr2[:]
	fmt.Println("This is s3:=arr2[2:] ")
	fmt.Println(s3)

	updateSlice(s1)
	fmt.Println("after update s1 ")
	fmt.Println("s1:")
	fmt.Println(s1)
	fmt.Println("arr2:")
	fmt.Println(arr2)

	updateSlice(s2)
	fmt.Println("after update s2 ")
	fmt.Println("s2:")
	fmt.Println(s2)
	fmt.Println("arr2:")
	fmt.Println(arr2)

	// reslice
	// s2:={2,3,4,5,6,7}
	s4 := s2[2:4]
	// s4:={4,5}, can be extended to include 6 and 7
	fmt.Println("after reslice s4")
	fmt.Println(s4)
	updateSlice(s4)
	fmt.Println(s4)
	fmt.Println(s2)
	fmt.Println(arr2)

	arr2 = [...]int{0,1,2,3,4,5,6,7}
	s2 = arr2[2:]
	// s2:={2,3,4,5,6,7}
	fmt.Println("s2 len:%v, cap:%v",len(s2),cap(s2))

	s4 = s2[2:4]
	// s4:={4,5}, can be extended to include 6 and 7
	fmt.Println("s4 len:%v, cap:%v",len(s4),cap(s4))

	s5 := s4[:3] // index 3 is not included
	// s5:={4,5,6}
	fmt.Println("s5 len:%v, cap:%v",len(s5),cap(s5))

	fmt.Println(s5)
	// slice can be extended forward to tail but not backward to head


	// append passes value, so we should get the value
	s4 = append(s4,10)
	// s4:={4,5,10}
	// arr2:=[0 1 2 3 4 5 10 7]
	fmt.Println(s4)
	fmt.Println(arr2)
	// sometimes append will change ptr and cap. Then this slice will point to a new array(created in memory)
	// rather than the original array. like if we append some elems to s2, the value in corresponding index in arr2 will be changed
	// But not the case if the append position is larger than arr2's len in s3, s4, s5
	s4 = append(s4,10)
	fmt.Println(s4)
	fmt.Println(arr2)
	// s4:={4,5,10,10}
	// arr2:=[0 1 2 3 4 5 10 10]

	s4 = append(s4,10)
	fmt.Println(s4)
	fmt.Println(arr2)
	// s4:={4,5,10,10,10}
	// arr2:=[0 1 2 3 4 5 10 10](unchanged)

	// slice initialization
	// s6 knows len but not initial values
	s6 := make([]int, 4)
	fmt.Println(s6)
	//s6 knows cap but not initial values
	s7 := make([]int, 10, 32)
	fmt.Println(s7)

	// delete from slice
	s4 = append(s4[:2],s6[3:]...)
	fmt.Println(s4)

	m := map[string][]int {
		"name": s4,
		"okay": s5,
	}

	mm := make(map[string]int)
	fmt.Println(mm)

	// delete elem in map
	delete(m,"name")

	if val, ok := m["name"]; ok {
		fmt.Println(val)
	} else {
		fmt.Println("not exists")
	}

	for k, v := range m{
		// the order is not guaranteed
		fmt.Println(k, v)
	}

	// besides slice, map, function, all types, even Struct, can be key
}


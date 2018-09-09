package main

import "fmt"

func maxSubArrayWithoutRepeatedChar(s string) int {
	// [l,r] without repeated char
	if s == "" {
		return 0
	}
	l := 0
	r := 0
	m := make(map[rune]bool)
	maxLength := 1

	arr := []rune(s)

	m[arr[l]] = true
	r++

	for r < len(arr) {
		if _, ok := m[arr[r]]; l < r && ok {
			if r-l > maxLength {
				maxLength = r - l
			}
			delete(m, arr[l])
			l++

		} else {
			m[arr[r]] = true
			r++
			if r-l > maxLength {
				maxLength = r - l
			}
		}

	}
	return maxLength
}
func main() {

	m := map[string]bool{
		"name": false,
		"okay": true,
	}

	mm := make(map[string]int)
	fmt.Println(mm)

	// delete elem in map
	delete(m, "name")

	if val, ok := m["name"]; ok {
		fmt.Println(val)
	} else {
		fmt.Println("not exists")
	}

	for k, v := range m {
		// the order is not guaranteed
		fmt.Println(k, v)
	}

	// in addition to slice, map, function, all other types, even Struct, can be key

	fmt.Println(maxSubArrayWithoutRepeatedChar("abcabcbb"))
	fmt.Println(maxSubArrayWithoutRepeatedChar("bbbbb"))
	fmt.Println(maxSubArrayWithoutRepeatedChar(""))
	fmt.Println(maxSubArrayWithoutRepeatedChar("pwwkew"))
	fmt.Println(maxSubArrayWithoutRepeatedChar("b"))
	fmt.Println(maxSubArrayWithoutRepeatedChar("abcdefg"))
	fmt.Println(maxSubArrayWithoutRepeatedChar("好呀好"))
}

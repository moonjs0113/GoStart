package main

import "fmt"

func Condtions() {
	ifSign := 10
	if ifSign > 0 {
		fmt.Println("if 문")
	}

	if ifSign < 0 {
		fmt.Println("if 문")
	} else {
		fmt.Println("else 문 ")
		fmt.Println("} else { 이렇게 안쓰고")
		fmt.Println("}\nelse { 이렇게 쓰면 문법 에러")
		fmt.Println("else if 도 당연? 가능")
	}

	var switchChar = "a"

	switch switchChar {
	case "s":
		fmt.Println("Switch s")
	case "a", "A", "ㅁ":
		fmt.Println("Switch a Default 없어도 작동")
	case "F":
		fmt.Println("Switch F")
		//default:
		//	fmt.Println("Switch Default")
	}
}

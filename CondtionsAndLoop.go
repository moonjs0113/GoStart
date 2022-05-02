package main

import "fmt"

func CondtionsAndLoop() {
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
	for i := 0; i < 5; i++ {
		fmt.Println("반복은 i:=0; i < 5; i++ 처럼")
		if i == 1 {
			fmt.Println("continue 키워드")
			continue
		}
		fmt.Printf("C, Java랑 굉장히 비슷하다. %d\n\n", i)
	}

	arrayInt := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for idx, value := range arrayInt {
		fmt.Println("for idx, value := range arrayInt")
		fmt.Printf("이렇게도 된다. idx: %v, value: %v\n", idx, value*2)
	}
}

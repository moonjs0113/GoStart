package main

import "fmt"

func Functions() {
	parameterFunc("I am String Value")
	sumResult := sumFunc(1, 4)
	fmt.Println("language func 예제 국룰:", sumResult)
	sumResult = sumFuncReturnName(4, 7)
	fmt.Println(sumResult)
	multiInt, multiString, _ := multiReturnFunc(3, 6)
	fmt.Println(multiInt, multiString)
	recursiveFunc(0)
}

func parameterFunc(name string) {
	fmt.Println("parameter value:", name)
	fmt.Println("func 이름(변수 타입, 변수 타입,...)", name)
}

func sumFunc(x int, y int) int {
	return x + y
}

// 이건 조금 신기한 문법
func sumFuncReturnName(x int, y int) (result int) {
	result = x + y
	return // result 써도 되고 안 써도 된다!
}

// 이건 조금 신기한 문법
func multiReturnFunc(x int, y int) (add int, addString string, double int) {
	add = x + y
	addString = "Oh!!!!"
	double = 2 * (x + y)
	return // 값이 한번에 리턴된다?
}

func recursiveFunc(x int) int {
	if x == 11 {
		return 0
	}
	fmt.Println(x)
	return recursiveFunc(x + 1)
}

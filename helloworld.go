// Package 선언
package main

import "fmt"

func main() {
	//Variable()
	//Output()
	//Array()
	//CondtionsAndLoop()
	Functions()
	Structures()
	Maps()
}

func Variable() {
	// Go Syntax
	fmt.Println("Hello Go World, Write in Go")

	// Go Variables

	// Declare Variables
	var stringType string = "String Type"
	var stringType2 = "타입 추론"
	// var stringType string 안됨
	intType := 123 // 타입 추론
	float32Type := 123.567
	boolType := true
	// var intType := 123 안됨
	intType = 234
	fmt.Println(stringType)
	fmt.Println(stringType2)
	fmt.Println(intType)
	fmt.Println(float32Type)
	fmt.Println(boolType)
	fmt.Println("--------------------")

	var initString string
	var initInt int
	var initFloat32 float32
	var initBool bool

	fmt.Println("선언만 하고 값 할당을 안 해도 된다.")
	fmt.Println(initString)
	fmt.Println(initInt)
	fmt.Println(initFloat32)
	fmt.Println(initBool)
	fmt.Println("--------------------")

	// var vs :=
	// var 는 선언과 할당이 자유롭다.
	// 코드블럭 밖에서 선언된 변수라도 사용이 가능하다.

	// := 는 코드블럭(func) 안에서 사용할 변수일 때만 사용한 선언 기호
	// 따라서 코드블럭 밖에서는 선언 자체가 불가능하다.
	// 선언과 할당이 받드시 같이 이루어진다.

	// Go Multiple Variable Declaration
	fmt.Println("한줄에 여러 변수를 선언하거나, 그룹으로 선언이 가능하다.")
	var multipleA, multipleB, multipleC int = 1, 3, 5
	var multipleIntA, multipleStringB = 6, "Hello"
	multipleIntC, multipleStringD := 7, "World!"

	var (
		groupIntB    int = 1
		groupIntA    int
		groupStringC string = "Group Declaration String C"
	)
	fmt.Println(multipleA, multipleB, multipleC)
	fmt.Println(multipleIntA, multipleStringB)
	fmt.Println(multipleIntC, multipleStringD)
	fmt.Println(groupIntB, groupIntA, groupStringC)
	fmt.Println("--------------------")

	fmt.Println("const 는 상수다.")
	const constInt = 0
	const constString string = "Const String"
	fmt.Println(constInt, constString)
	fmt.Println("--------------------")

}

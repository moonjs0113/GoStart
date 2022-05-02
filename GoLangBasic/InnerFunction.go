package GoLangBasic

import "fmt"

func InnerFunc() {
	deferFunction()
	println(panicFunction())
}

func deferFunction() {
	println("defer는 특정 문장이나 함수가 defer를 감싸는 함수가 리턴되지 직전에 실행되도록 한다.")
	defer println("짜잔")
	println("함수 문 닫습니다~")
}

func panicFunction() int {
	println("panic은 현재 함수를 멈추고 defer 코드를 즉시 실행 시킨다.")
	a := 10
	println(a)
	defer print("panic 속 defer")
	defer recoverFunction()
	panic(a)

	a = 40
	return a
}

func recoverFunction() {
	if r := recover(); r != nil {
		println(r)
		fmt.Println(r)
	}
}

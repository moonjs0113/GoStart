package W3School

import "fmt"

func Output() {
	var j int = 15

	fmt.Println("print 는 현재 커서 위치에 단순 프린트")
	fmt.Println("println 은 문자열 끝에 개행문자 추가된 프린트")
	fmt.Println("printf 은 포맷이 추가된 프린트, 자동 개행은 안된다.")
	fmt.Printf("j has value: %v and type: %T\n", j, j)
	fmt.Println("%v: Value, %T: Type")
	fmt.Println("c랑 java 포맷 문자랑 비슷하다!")
	fmt.Println("%d: int, %f: float, %s: string, %t: bool")
	fmt.Println("--------------------")
}

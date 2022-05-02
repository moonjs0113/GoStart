package GoLangBasic

import "fmt"

type Rect struct {
	width, height int
}

func (rect Rect) area() int {
	println("Struct Method를 밖에서 만든다.")
	println("func (receiver 구조체) 메소드명() 반환타입")
	return rect.width * rect.height
}

func (rect *Rect) pointerArea() int {
	println("Pointer Method는 레퍼런스주소로 값을 가져온다.")
	println("-> Method 내에서 값을 변경하면 바뀐다.")
	rect.height *= 10
	return rect.width * rect.height
}

func StructMethod() {
	rect := Rect{height: 10, width: 55}
	area := rect.area()
	println(area)
	fmt.Println(rect)
	pointerArea := rect.pointerArea()
	println(pointerArea)
	fmt.Println(rect)
}

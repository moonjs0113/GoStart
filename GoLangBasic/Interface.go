package GoLangBasic

import (
	"fmt"
	"math"
)

// Interface
type Shape interface {
	area() float64
	perimeter() float64
}

// Struct
type InterfaceRect struct {
	width, height float64
}
type InterfaceCircle struct {
	radius float64
}

func (r InterfaceRect) area() float64 {
	return r.width * r.height
}

func (r InterfaceRect) perimeter() float64 {
	return 2 * (r.width + r.height)
}

func (c InterfaceCircle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c InterfaceCircle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func Interface() {
	println("Interface는 Method의 집합이다.")
	println("java의 interface랑 비슷하고, Swift의 Protocol이랑 비슷하다.")
	rect := InterfaceRect{10.0, 25.0}
	cir := InterfaceCircle{15}
	showArea(rect, cir)

	println("빈 Interface는 Any Type을 의미한다.")
	var emptyInterface interface{}
	emptyInterface = 1     // 가능
	emptyInterface = "Tom" // 가능
	stringValue := emptyInterface.(string)
	println(emptyInterface)     // 제대로 안 나온다.
	fmt.Println(emptyInterface) // fmt에서 찍어야 제대로 나온다.
	println(stringValue)
}

func showArea(shapes ...Shape) {
	for _, s := range shapes {
		a := s.area()
		p := s.perimeter()
		println(a, p)
	}
}

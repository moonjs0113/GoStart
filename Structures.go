package main

import "fmt"

func Structures() {
	var person0 Person
	var person1 = Person{
		name:      "Rey",
		age:       100,
		subscript: "객체 예제 국룰"}

	person0.name = "Good"
	person0.age = 222
	person0.subscript = "abcdefghijklmnopqrstuwxyz"

	fmt.Println(person0.age)
	fmt.Println(person1.name)
}

type Person struct {
	name      string
	age       int
	subscript string
}

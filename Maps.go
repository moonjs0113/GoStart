package main

import "fmt"

func Maps() {
	var mapValue = map[string]string{"brand": "Ford", "model": "Mustang", "year": "1964"}

	fmt.Println("Dictionary랑 비슷하다. -> 순서 보장 X")
	fmt.Println("map[keyType]valueType{:,:,...}")
	fmt.Println(mapValue)
	delete(mapValue, "brand")
	fmt.Println(mapValue)

	fmt.Println("make로도 만들 수 있다.")
	var makeMap = make(map[string]string)
	makeMap["a"] = "A"
	fmt.Println("make(map[string]string)")

	var notNilMap = make(map[string]string)
	var nilMap map[string]string

	fmt.Println("is nil? ", notNilMap == nil)
	fmt.Println("var notNilMap = make(map[string]string)")
	fmt.Println("is nil? ", nilMap == nil)
	fmt.Println("var nilMap map[string]string")

	value0, exist0 := mapValue["model"]
	value1, exist1 := mapValue["price"]
	fmt.Println(value0, exist0)
	fmt.Println(value1, exist1)

	fmt.Println("레퍼런스 타입이라 다른 변수에 저장 후 값을 변경하면 같이 바뀐다.")
	temp := mapValue
	temp["price"] = "So Expensive"
	fmt.Println(mapValue)

	for key, value := range mapValue {
		fmt.Println(key, value)
	}
}

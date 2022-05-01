package main

import "fmt"

func Array() {
	var array0 = [5]int{4, 5, 6, 7, 8}
	array1 := [...]int{4, 5, 6, 7, 8, 9, 10}
	array2 := [6]int{123, 456}
	array3 := [6]string{2: "Hello", 4: "World", 5: "String"}
	fmt.Println("Array [길이]타입{값들} 로 선언")
	fmt.Println("길이를 ...로 주면 길이를 추론한다.")
	array0[0] = 123
	fmt.Println(array0)
	fmt.Println(array1)
	fmt.Println(array2)
	fmt.Println(array3)
	fmt.Println(len(array3))
	fmt.Println("--------------------")

	fmt.Println("Slice라는 개념이 있다.")
	sliceInt := []int{}
	arr1 := [6]int{10, 11, 12, 13, 14, 15}
	myslice := arr1[2:3]
	fmt.Println(sliceInt)
	fmt.Println("length와 capacity가 있다.")
	fmt.Println("length는 현재 길이, capacity는 append 등으로 늘어날 수 있는 최대 길이다.")
	fmt.Printf("myslice = %v\n", myslice)
	fmt.Printf("length = %d\n", len(myslice))
	fmt.Printf("capacity = %d\n", cap(myslice))

	fmt.Println("make([]type, length, capacity) 로도 slice 생성이 가능한다.")
	makeSlice := make([]int, 4, 5)
	fmt.Println(makeSlice)
	makeSlice[0] = 234
	fmt.Println(makeSlice)
	fmt.Println(len(makeSlice))
	fmt.Println(cap(makeSlice))
	makeSlice = append(makeSlice, 20, 21)
	fmt.Println(makeSlice)
	fmt.Println(len(makeSlice))
	fmt.Println(cap(makeSlice))
	fmt.Println("append로 Capacity가 초과되면, Capacity가 두배인 Slice를 생성하여 저장한다.")
	fmt.Println("배열에서 필요한 부분만 Slice로 만들고 싶다면, Copy를 쓰는게 좋다.")
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	// 필요한 만큼 잘라서 Slice 생성
	neededNumbers := numbers[:len(numbers)-10]
	// 자른 Slice의 길이를 이용해 다시 Init Slice 생성
	numbersCopy := make([]int, len(neededNumbers))
	// 값 복사
	copy(numbersCopy, neededNumbers)

	fmt.Printf("neededNumbers = %v\n", neededNumbers)
	fmt.Printf("length = %d\n", len(neededNumbers))
	fmt.Printf("capacity = %d\n", cap(neededNumbers))

	fmt.Printf("numbersCopy = %v\n", numbersCopy)
	fmt.Printf("length = %d\n", len(numbersCopy))
	fmt.Printf("capacity = %d\n", cap(numbersCopy))
	fmt.Println("--------------------")
}

package GoLangBasic

func gotoKeyword() {
	var a = 1
	for a < 15 {
		if a == 5 {
			a += a
			continue // for루프 시작으로
		}
		a++
		if a > 10 {
			break //루프 빠져나옴
		}
	}
	if a == 11 {
		goto END //goto 사용예
	}
	println(a)

END:
	println("End")

	i := 0
L1:
	for {
		if i == 0 {
			break L1
		}
	}
	println("Break Label")

	multiParameterFunction("a", "b", "banana", "piano")
	anonymousFunction()
	closureExample()
}

func multiParameterFunction(strings ...string) {
	for index, string := range strings {
		println(index, string)
	}
}

type typealias func(int, int) int

func anonymousFunction() {
	sum := func(n ...int) int {
		s := 0
		for _, i := range n {
			s += i
		}
		return s
	}

	result := sum(1, 2, 3, 4, 5)
	println(result)

	var firstClassFunction func(typealias) int
	firstClassFunction = functionParameter
	println(firstClassFunction(paramSumFunc))
}

func functionParameter(sumFunc typealias) int {
	result := sumFunc(3, 7)
	return result
}

func paramSumFunc(x int, y int) int {
	return x + y
}

func closure() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func closureExample() {
	next := closure()

	println(next()) // 1
	println(next()) // 2
	println(next()) // 3

	anotherNext := closure()
	println(anotherNext()) // 1 다시 시작
	println(anotherNext()) // 2
}

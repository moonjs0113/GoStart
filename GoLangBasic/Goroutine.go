package GoLangBasic

import (
	"fmt"
	"sync"
	"time"
)

func Goroutine() {
	asyncFunction()
	anonymousFunctionAsync()
}

func say(s string) {
	for i := 0; i < 10; i++ {
		fmt.Println(s, "***", i)
	}
}

func asyncFunction() {
	say("Sync")

	go say("Async1")
	go say("Async2")
	go say("Async3")

	// 3s Wait
	time.Sleep(time.Second * 1)
}

func anonymousFunctionAsync() {
	// WaitGroup 생성. 2개의 Go루틴을 기다림.
	var wait sync.WaitGroup
	wait.Add(2)

	// 익명함수를 사용한 goroutine
	go func() {
		defer wait.Done() //끝나면 .Done() 호출
		fmt.Println("Hello")
	}()

	// 익명함수에 파라미터 전달
	go func(msg string) {
		defer wait.Done() //끝나면 .Done() 호출
		fmt.Println(msg)
	}("Hi")

	wait.Wait() //Go루틴 모두 끝날 때까지 대기
}

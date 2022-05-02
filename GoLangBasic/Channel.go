package GoLangBasic

import (
	"fmt"
	"time"
)

func Channel() {
	channel_int()
	channel_anonymousFunc()
	channel_buffer()
	channel_parameter()
	channel_close()
	channel_select()
}

func channel_int() {
	ch := make(chan int)

	go func() {
		ch <- 123
	}()

	var i int
	i = <-ch
	println(i)
}

func channel_anonymousFunc() {
	done := make(chan bool)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(i)
		}
		done <- true
	}()
	boolean := <-done
	println(boolean)
}

func channel_buffer() {
	//c := make(chan int)
	// 수신루틴이 없으므로 데드락
	//c <- 1
	// 코멘트해도 데드락 (별도의 Go루틴없기 때문)
	//fmt.Println(<-c)

	ch := make(chan int, 1)
	ch <- 112233
	fmt.Println(<-ch)
}

func channel_parameter() {
	ch := make(chan string, 1)
	sendChan(ch)
	receiveChan(ch)
	twoWayChan(ch)
}

// 받으려는 채널로 넘김
func sendChan(ch chan<- string) {
	ch <- "Data"
}

// 주려는 채널로 넘김
func receiveChan(ch <-chan string) {
	data := <-ch
	fmt.Println(data)
}

func twoWayChan(ch chan string) {
	ch <- "Two-Way Channel Parameter"
	data := <-ch
	fmt.Println(data)
}

func channel_close() {
	ch := make(chan int, 2)

	ch <- 1
	ch <- 2

	close(ch)

	// 닫아도 수진은 된다.
	println(<-ch)
	println(<-ch)

	// 채널 버퍼에 데이터가 있는지 확인
	if _, success := <-ch; !success {
		println("더이상 데이터 없음.")
	}

	// range로 확인해도 된다.
	// channel에 남아있는 데이터가 있으면 반복이 되고 아니면 멈춘다.
	for i := range ch {
		println(i)
	}
}

func channel_select() {
	done1 := make(chan bool)
	done2 := make(chan bool)

	go run1(done1)
	go run2(done2)

EXIT:
	for {
		select {
		case <-done1:
			println("run1 완료")

		case <-done2:
			println("run2 완료")
			break EXIT
			//default:
			//	println("wait done channel")
		}
	}
}

func run1(done chan bool) {
	time.Sleep(1 * time.Second)
	done <- true
}

func run2(done chan bool) {
	time.Sleep(2 * time.Second)
	done <- true
}

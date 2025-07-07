package main

import (
	"fmt"
	"time"
)

// fastSender 함수는 int 타입의 채널에 1부터 5까지의 정수를 0.1초 간격으로 보냅니다.
func fastSender(ch chan int) {
	defer close(ch) // 함수 종료 시 채널을 닫습니다.
	for i := 1; i <= 5; i++ {
		ch <- i
		time.Sleep(100 * time.Millisecond) // 0.1초 대기
	}
}

func slowSender(ch chan string) {
	defer close(ch) // 함수 종료 시 채널을 닫습니다.
	messages := []string{"hello", "world", "golang", "concurrency"}
	for _, msg := range messages {
		ch <- msg
		time.Sleep(1 * time.Second) // 1초 대기
	}
}

func main() {
	ints := make(chan int)
	go fastSender(ints)

	strings := make(chan string)
	go slowSender(strings)

	// 무한 루프를 돌면서 채널에서 데이터를 받습니다.
	// 두 채널 모두 닫히고 모든 데이터가 처리되면 루프를 종료합니다.
	for {
		select {
		// strings 채널에서 값을 받습니다. ok는 채널이 열려있으면 true, 닫혔으면 false입니다.
		case s, ok := <-strings:
			if !ok {
				// 채널이 닫혔으면 해당 case가 더 이상 선택되지 않도록 채널을 nil로 설정합니다.
				strings = nil
				// 두 채널 모두 nil이 되었는지 확인하여 main 루프를 종료합니다.
				if ints == nil {
					fmt.Println("모든 채널이 닫혔습니다. 프로그램 종료.")
					return // main 함수 종료
				}
				continue // 다음 select 반복으로 넘어갑니다.
			}
			fmt.Println("Received a string:", s)
		// ints 채널에서 값을 받음. ok는 채널이 열려있으면 true, 닫혔으면 false
		case i, ok := <-ints:
			if !ok {
				// 채널이 닫혔으면 해당 case가 더 이상 선택되지 않도록 채널을 nil로 설정
				ints = nil
				// 두 채널 모두 nil이 되었는지 확인해 main 루프를 종료
				if strings == nil {
					fmt.Println("모든 채널이 닫혔습니다. 프로그램 종료.")
					return // main 함수 종료
				}
				continue // 다음 select 반복으로 넘어감
			}
			fmt.Println("Received an int:", i)
		}
	}
}

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	joe := boring("joe")
	ann := boring("ann")

	for i := 0; i < 5; i++ {
		fmt.Println(<-joe)
		fmt.Println(<-ann)
	}

	fmt.Println("you are both boring; I'm leaving")
}

func boring(msg string) <-chan string {
	// chan string 은 리턴값임 참고로 <- means [recieve only], This function return channel but can only receive

	c := make(chan string)

	go func() { // 반환값이 없는 함수

		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}

	}() // 클로져를 형성하여 msg 와 c 에 접근

	return c // return the channel to the caller
}

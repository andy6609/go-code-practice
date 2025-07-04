package main

import (
	"fmt"
)

// adder function returns closure
// Closure remember variable sum, 그리고 호출 될 떄마다 값 호출

func adder() func(int) int { // 아무 인자 받지 않고, return 값 으로 "func(int) int 을 반환"
	sum := 0                 // adder 내부에서 선언된 local variable 이지만 그 안의 내부 func 에서도 계속 사용됨!!
	return func(x int) int { //adder 함수 내부에서 만들어지는 annonymous func
		sum += x
		return sum // Go 에서는 closure 선언 같은거 없이, 내부함수가 외부함수의 변수를 사용하면 자동 closure
	}
}

func main() {
	positive, negative := adder(), adder() // Each closure 생성(create)
	// 그리고 pos, neg 얘네들은 함수 자체를 담고(contain) 있는 변수임
	// 여기에는 adder()함수가 돌려준 '함수' 그 자체가 할당됨(allocated).
	for i := 0; i < 10; i++ {
		fmt.Println(positive(i), negative(-i))
	}
}

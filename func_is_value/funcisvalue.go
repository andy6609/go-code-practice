package main

import (
	"fmt"
	"math"
)

// compute 함수는 "float54 두 개를 받아서 float64를 반환 하는 함수"
// 인자로 받고 returen 값으로 float64를 내보냄
func compute(fn func(float64, float64)float64) float64{
	return fn(3,4)
}
// fn 없이 그냥 func 쓰면 어떤 함수가 들어오는지 Go가 식별 x

func main(){
	hypot := func(x,y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}

	// hypot 직접 실행 
	fmt.Println(hypot(5,12)) // 13

	// hypot 함수 자체를 compute 함수에 전달 
	fmt.Println(compute(hypot)) // 5(sqrt(3*3 + 4*4) = 5)

	// math.Pow 함수도 인자로 전달 가능 
	fmt.Println(compute(math.Pow)) //
}
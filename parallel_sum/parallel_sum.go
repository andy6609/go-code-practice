package main

import "fmt"

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v // v 는 s의 각 요소의 값(예 1,2,3) 각각 다 take it out and add to sum
	}
	c <- sum
} // After calculating sum of numbers, func sent it to channel not returning

func main() {
	s := []int{7, 2, 8, -9, 4, 0}

	mid := len(s) / 2

	leftSum := make(chan int) // what if we don't make (chan int)
	rightSum := make(chan int)

	go sum(s[:mid], leftSum) // Start two new goroutines. Paralel
	go sum(s[mid:], rightSum)

	x := <-leftSum // why not x <- leftSum?
	y := <-rightSum

	fmt.Println(x, y, x+y)
}

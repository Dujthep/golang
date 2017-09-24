package main

import (
	"fmt"
	"math/cmplx"
)

var (
	ToBe bool = false
	MaxInt uint64 = 1<<64 - 1
	z complex128 = cmplx.Sqrt(-5 + 12i)
	arr = []string{"a", "b", "c", "d", "e", "f"}
)

const (
	
	Pi = 3.14
	a = iota + 5
	b
	c
	d
)

func add(x int, y int) int {
	return x + y
}

func loopTest() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum +=i
		fmt.Println(sum)
	}
}

func loopTest2() {
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}

func sliceArray() {
	tmp := arr[1:3]
	fmt.Println(tmp)
}

func appendArray() {
	a := []int{}
	for i := 0; i < 10; i++ {
		a = append(a, i)
	}
	fmt.Println("length = %d cap = %d %v\n",len(a), cap(a), a)
}

func main() {
	// loopTest()
	// loopTest2()
	// slice()
	appendArray()
}

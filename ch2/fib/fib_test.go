package fib

import (
	"fmt"
	"testing"
)

// 斐波拉切数列
func TestFibList(t *testing.T) {
	// var a int = 1
	// var b int = 1
	var (
		a = 1
		b = 1
	)
	fmt.Print(a)
	for i := 0; i < 5; i++ {
		fmt.Print(" ", b)
		tmp := a
		a = b
		b = tmp + a
	}
	fmt.Println()
}

// 交换两个变量的值
func TestExchange(t *testing.T) {
	// 常规
	a, b := 1, 2
	tmp := a
	a = b
	b = tmp
	fmt.Print(a, b)
	fmt.Println()

	// 简洁做法
	c, d := 3, 4
	c, d = d, c
	fmt.Print(c, d)
}

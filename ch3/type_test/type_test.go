package type_test

import (
	"fmt"
	"testing"
)

func TestImplicit(t *testing.T) {
	var a int = 1
	var b int64
	// b = a // Error, 不能进行隐式类型转换， 别名也不行
	b = int64(a) // 需要显式类型转换
	fmt.Println(a, b)
}

func TestPoint(t *testing.T) {
	a := 1
	aPtr := &a
	fmt.Println(a, aPtr)
	fmt.Printf("%T %T", a, aPtr)

	// aPtr = aPtr + 1  报错，不支持指针运算
}

func TestString(t *testing.T) {
	var s string
	fmt.Println("*" + s + "*") // string的默认值是空值，加上 * 表示
	fmt.Println(len(s))        // 通过 len 函数打印出长度，0
}

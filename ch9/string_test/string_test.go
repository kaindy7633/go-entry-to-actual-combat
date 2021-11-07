package stringtest

import (
	"fmt"
	"testing"
)

func TestString(t *testing.T) {
	var s string
	fmt.Println(s)
	s = "Hello"
	fmt.Println(len(s))

	s = "\xE4\xB8\xA5"
	fmt.Println(s)
	fmt.Printf("中文字符长度：%d\n", len(s))

	s = "中"
	fmt.Printf("一个中文字符的长度：%d\n", len(s))
	c := []rune(s)
	fmt.Printf("长度：%d\n", len(c))
	fmt.Printf("中 unicode %x\n", c[0])
	fmt.Printf("中 UTF8 %x\n", s)
}

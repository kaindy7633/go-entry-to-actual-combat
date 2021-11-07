package constant

import (
	"fmt"
	"testing"
)

const (
	Monday = iota + 1
	Tuesday
	Wednesday
)

const (
	Readable = 1 << iota
	Writable
	Executable
)

func TestConstant(t *testing.T) {
	fmt.Println(Monday, Tuesday, Wednesday)
}

func TestConstantByte(t *testing.T) {
	fmt.Println(Readable)
	fmt.Println(Writable)
	fmt.Println(Executable)
}

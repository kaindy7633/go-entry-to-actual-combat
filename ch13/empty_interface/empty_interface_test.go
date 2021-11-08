package empty_interface

import (
	"fmt"
	"testing"
)

func DoSomething(p interface{}) {
	if i, ok := p.(int); ok {
		fmt.Println("Integer", i)
		return
	}

	if s, ok := p.(string); ok {
		fmt.Println("string", s)
		return
	}

	fmt.Println("Unknow type")
}

// switch 写法
func DoSomethingUseSwitch(p interface{}) {
	switch v := p.(type) {
	case int:
		fmt.Println("Integer", v)
	case string:
		fmt.Println("String", v)
	default:
		fmt.Println("Unknow type")
	}
}

func TestEmptyInterfaceAssertion(t *testing.T) {
	DoSomething(10)
	DoSomething("10")
}

func TestEmptyInterfaceAssertion2(t *testing.T) {
	DoSomethingUseSwitch(10)
	DoSomethingUseSwitch("10")
}

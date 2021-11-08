package encap

import (
	"fmt"
	"testing"
)

type Employee struct {
	Id   int
	Name string
	Age  int
}

func (e Employee) String() string {
	return fmt.Sprintf("ID:%d-Name:%s-Age:%d", e.Id, e.Name, e.Age)
}

func (e *Employee) StringByPoint() string {
	// 推荐使用传入指针的方式
	return fmt.Sprintf("ID:%d-Name:%s-Age:%d", e.Id, e.Name, e.Age)
}

func TestCreateEmployeeObj(t *testing.T) {
	e := Employee{0, "Bob", 20}
	e1 := Employee{Id: 1, Name: "Mike", Age: 30}
	e2 := new(Employee)
	e2.Id = 1
	e2.Name = "Rose"
	e2.Age = 22

	fmt.Println(e)
	fmt.Println(e1)
	fmt.Println(e1.Id)
	fmt.Println(e2)
	fmt.Printf("e is %T\n", e)
	fmt.Printf("e2 is %T\n", e2)
}

func TestStructOperations(t *testing.T) {
	e := Employee{0, "Bob", 20}
	fmt.Println(e.String())

	ePtr := Employee{1, "Mike", 30}
	fmt.Println(ePtr.StringByPoint())
}

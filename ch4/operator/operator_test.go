package operator

import (
	"fmt"
	"testing"
)

func TestCompareArray(t *testing.T) {
	a := [...]int{1, 2, 3, 4}
	b := [...]int{1, 3, 4, 5}
	c := [...]int{1, 2, 3, 4, 5}
	d := [...]int{1, 2, 3, 4}

	fmt.Println(a == b)
	fmt.Println(a == d)
	// fmt.Println(b == c)  维度不同，无法进行比较
	fmt.Println(c)
}

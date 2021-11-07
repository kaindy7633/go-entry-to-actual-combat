package array

import (
	"fmt"
	"testing"
)

func TestArrayInit(t *testing.T) {
	var arr [3]int
	arr1 := [4]int{1, 2, 3, 4}
	arr2 := [...]int{1, 3, 4}
	fmt.Println(arr[1], arr[2])
	fmt.Println(arr1, arr2)
}

func TestArrayTravel(t *testing.T) {
	arr3 := [...]int{1, 2, 3, 4}
	for i := 0; i < len(arr3); i++ {
		fmt.Println(arr3[i])
	}

	for idx, e := range arr3 {
		fmt.Println(idx, e)
	}
}

func TestArraySection(t *testing.T) {
	arr4 := [...]int{1, 2, 3, 4, 5}
	fmt.Println(arr4[:3])
}

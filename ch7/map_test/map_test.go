package map_tset

import (
	"fmt"
	"testing"
)

func TestMapInit(t *testing.T) {
	m1 := map[int]int{1: 1, 2: 4, 3: 9}
	fmt.Println(m1)
	fmt.Println(m1[2])
	fmt.Println("len m1", len(m1))

	m2 := map[int]int{}
	m2[4] = 16
	fmt.Println(m2)
}

func TestAccessNotExistingKey(t *testing.T) {
	m1 := map[int]int{}
	if v, ok := m1[3]; ok {
		fmt.Println(v)
	} else {
		fmt.Println("not exists")
	}
}

func TestTravelMap(t *testing.T) {
	m1 := map[int]int{1: 1, 2: 4, 3: 9}
	for k, v := range m1 {
		fmt.Printf("Key: %d, Val: %d\n", k, v)
	}
}

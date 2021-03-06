package slice

import (
	"fmt"
	"testing"
)

func TestSliceInit(t *testing.T) {
	var s0 []int
	fmt.Printf("长度：%d， 容量：%d\n", len(s0), cap(s0))

	// 添加元素
	s0 = append(s0, 1)
	fmt.Printf("长度：%d， 容量：%d\n", len(s0), cap(s0))

	// 带值初始化
	s1 := []int{1, 2, 3, 4}
	fmt.Printf("长度：%d， 容量：%d\n", len(s1), cap(s1))

	s2 := make([]int, 3, 5)
	fmt.Printf("长度：%d， 容量：%d\n", len(s2), cap(s2))
}

func TestSliceGrowing(t *testing.T) {
	s := []int{}
	for i := 0; i < 10; i++ {
		s = append(s, i)
		fmt.Printf("长度：%d， 容量：%d\n", len(s), cap(s))
	}
}

func TestSliceShareMemory(t *testing.T) {
	year := []string{"Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"}
	Q2 := year[3:6]
	fmt.Printf("长度：%d， 容量：%d\n", len(Q2), cap(Q2))

	summer := year[5:8]
	fmt.Printf("长度：%d， 容量：%d\n", len(summer), cap(summer))
	summer[0] = "Unknow"
	fmt.Println(Q2)
	fmt.Printf("长度：%d， 容量：%d\n", len(Q2), cap(Q2))
}

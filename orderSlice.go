// test
package main

import (
	"fmt"
	"time"
)

func main() {
	s1 := []int{1, 2, 3}
	s2 := []int{1, 2, 3}
	t1 := time.Now()
	for i := 0; i < 3; i++ {
		tmp := append([]int{}, s1[3:]...)
		s1 = append(s1[0:3], i)
		s1 = append(s1, tmp...)
		//s1 = append(s1[0:3], append([]int{i}, s1[3:]...)...)
	}
	fmt.Println(s1)
	t2 := time.Now()
	for i := 0; i < 3; i++ {
		//tmp := append([]int{}, s2[3:]...)
		s2 = append(s2, i)
		fmt.Println("b:", s2)
		fmt.Println("b:", s2[3:len(s2)], s2[2:len(s2)-1])
		copy(s2[3:len(s2)], s2[2:len(s2)-1])
		fmt.Println("m:", s2)
		s2[3] = i
		fmt.Println("l:", s2)
	}
	fmt.Println(s2)
	t3 := time.Now()
	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[i] {
			fmt.Println("error")
		}
	}
	fmt.Println(t3.Sub(t2), t2.Sub(t1))
}

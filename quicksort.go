// test
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	num := 5
	arr := make([]int, 0, num)
	rand.Seed(time.Now().Unix())
	s := time.Now()
	for i := 0; i < num; i++ {
		randNum := rand.Intn(100)
		arr = append(arr, randNum)
	}
	s1 := time.Now()
	fmt.Println("use time:", s1.Sub(s).Seconds())
	fmt.Println("out:", arr)
	quickSort(arr, 0, len(arr)-1)
	s2 := time.Now()
	fmt.Println("use time:", s2.Sub(s1).Seconds())
	//fmt.Println(arr)

	fmt.Println(7 | (1 << 20))
}
func quickSort(data ArraySlice, low, row int) {
	fmt.Println(low, row)
	if low < row {
		i, j, x, last := low, row, low, 0 //0就是使用第一个作为基准值,last这个变量时为了基准最后一次交换变量时出现在那次
		for i < j {
			fmt.Println("in data:", data)
			fmt.Println("in i=j:", i, j)
			for i < j && data.Less(x, j) { //比x小的放在前面出现的坑中
				j--
			}
			fmt.Println("in1 i=j:", i, j)
			if i < j {
				data.Swap(i, j)
				i++
				x = j
				last = 1
			}
			fmt.Println("in data:", data)
			fmt.Println("in i=j=x:", i, j, x)
			for i < j && data.Less(i, x) { //比x大的放在后面出现的坑中
				i++
			}
			fmt.Println("in2 i=j:", i, j)
			if i < j {
				data.Swap(i, j)
				j--
				x = i
				last = -1
			}
			fmt.Println("in data:", data)
			fmt.Println("in i=j=x:", i, j, x)
		}
		if last == 1 {
			data.Swap(j, x)
		} else if last == -1 {
			data.Swap(i, x)
		}
		quickSort(data, low, i-1)
		quickSort(data, i+1, row)
	}
}

type ArraySlice []int

func (s ArraySlice) Len() int {
	return len(s)
}
func (s ArraySlice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s ArraySlice) Less(i, j int) bool {
	if s[i] < s[j] {
		return true
	} else {
		return false
	}
}

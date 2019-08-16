package main

import "fmt"

func makeSlice () {
	// 16 是slice的长度  32 是cap的长度
	s := make([]int, 16, 32)
	fmt.Println(s)
}

func copySlice()  {
	arr := [...] int {1, 2, 3, 4, 5, 6, 7}
	s1 := arr[2:6]
	s2 := make([]int, 16, 32)
	copy(s2, s1)
	fmt.Println(s2)
}

func main() {
	arr := [...] int {1, 2, 3, 4, 5, 6, 7}
	s1 := arr[2:6]
	// s1 = [3, 4, 5, 6] 对应的 cap [3, 4, 5, 6, 7]
	// 下标越界，在不超过底层数组cap值还是能显示出来 如下  s2已经越界了
	s2 := s1[3:5]
	// s2 = [6, 7]
	fmt.Printf("s1 =%v,	len(s1) = %d,  cap(s1) = %d\n ", s1, len(s1), cap(s1))
	fmt.Println("s2: ", s2)

	// 向后添加元素 如果超过 cap的长度go自动会在底层创建一个更长的数组
	s3 := append(s2, 10)

	fmt.Println("s3: ", s3)
}

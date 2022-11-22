package main

import (
	"fmt"
	"sort"
)

func main() {

	//map有序输出
	m1 := make(map[int]string)
	m1[1] = "a"
	m1[2] = "b"
	m1[5] = "e"
	m1[4] = "d"
	m1[3] = "c"

	slic := []int{}
	for i := range m1 {
		slic = append(slic, i)
	}
	fmt.Println(slic)
	sort.Ints(slic)
	fmt.Println(slic)
	for i := 0; i < len(slic); i++ {
		fmt.Println(m1[slic[i]])
	}
}

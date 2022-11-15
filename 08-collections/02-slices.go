package main

import "fmt"

func main() {
	// var nos []int
	// var nos []int = []int{3, 1, 4, 2, 5}
	// var nos = []int{3,1,4,2,5}
	// nos := []int{3, 1, 4, 2, 5}
	// nos := []int{}
	var nos []int
	//nos[0] = 10
	nos = append(nos, 10)
	nos = append(nos, 20, 30, 40)

	hundreds := []int{100, 200, 300, 400, 500}
	nos = append(nos, hundreds...)
	fmt.Println(nos)

	/* x := []int{3, 1, 4, 2, 5}
	fmt.Println("Before sorting, x = ", x)
	sort(x)
	fmt.Println("After sorting, x = ", x) */

	//slicing
	/*
		[lo : hi] => from lo to hi-1
		[lo : ] => from lo to end of the slice
		[ : hi] => from 0 to hi-1
	*/

	x := nos[1:4]
	fmt.Println("nos[1:4] = ", x)
	x[0] = 999
	fmt.Println(x)
	fmt.Println(nos)
}

func sort(list []int) {
	for i := 0; i <= 3; i++ {
		for j := i + 1; j <= 4; j++ {
			if list[i] > list[j] {
				list[j], list[i] = list[i], list[j]
			}
		}
	}
}

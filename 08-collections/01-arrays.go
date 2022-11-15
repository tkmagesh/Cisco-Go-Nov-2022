package main

import "fmt"

func main() {
	//var nos [5]int
	//var nos [5]int = [5]int{3, 1, 4, 2, 5}
	//var nos = [5]int{3, 1, 4, 2, 5}
	//nos := [5]int{3, 1, 4, 2, 5}
	//nos := [5]int{}

	nos := [5]int{3, 1, 4, 2, 5}
	fmt.Println(nos)

	fmt.Println("Iterating an array (using indexer)")
	for i := 0; i < len(nos); i++ {
		fmt.Printf("nos[%d] = %d\n", i, nos[i])
	}

	fmt.Println("Iterating an array (using range)")
	/*
		for idx, no := range nos {
			fmt.Printf("nos[%d] = %d\n", idx, no)
		}
	*/
	for _, no := range nos {
		fmt.Printf("%d\n", no)
	}

	fmt.Println("Before sorting, nos = ", nos)
	sort(nos)
	fmt.Println("After sorting, nos = ", nos)

	var x *[5]int
	x = &nos

	//(*x)[0] = 100
	x[0] = 100
	fmt.Println(x, nos)

	nos1 := [5]int{10, 20, 30, 40, 50}
	nos2 := [5]int{10, 20, 30, 40, 50}
	fmt.Println(nos1 == nos2)

}

func sort(list [5]int) {
	for i := 0; i <= 3; i++ {
		for j := i + 1; j <= 4; j++ {
			if list[i] > list[j] {
				list[j], list[i] = list[i], list[j]
			}
		}
	}
}

package main

import "fmt"

func main() {
	//var productRanks map[string]int = make(map[string]int)
	/*
		productRanks := make(map[string]int)
		productRanks["Pen"] = 4
	*/
	// var productRanks map[string]int = map[string]int{}
	// var productRanks map[string]int = map[string]int{"Pen": 4, "Pencil": 2, "Marker": 5}
	var productRanks map[string]int = map[string]int{
		"Pen":    4,
		"Pencil": 2,
		"Marker": 5, //make sure the last item also has a comma (,)
	}
	productRanks["Notepad"] = 1
	fmt.Println(productRanks)
	fmt.Println("len(productRanks) =", len(productRanks))

	fmt.Println("Iterating a map")
	for key, value := range productRanks {
		fmt.Printf("productRanks[%q] = %d\n", key, value)
	}

	//check for the existence of a key
	keyToCheck := "Notepad"
	if rank, exists := productRanks[keyToCheck]; exists {
		fmt.Printf("Rank of %q is %d\n", keyToCheck, rank)
	} else {
		fmt.Printf("Key [%q] doesnot exist\n", keyToCheck)
	}

	keyToDelete := "Notepad"
	delete(productRanks, keyToDelete)
	fmt.Println(productRanks)
}

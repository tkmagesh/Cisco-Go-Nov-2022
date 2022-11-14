/* applied higher order functions */
package main

import (
	"fmt"
	"log"
)

func main() {

	logAdd := getLogOperation(add)
	result := logAdd(100, 200, 300, 400)
	fmt.Println("Result : ", result)

}

func getLogOperation(operation func(...int) int) func(...int) int {
	return func(nos ...int) int {
		log.Println("operation started")
		result := operation(nos...)
		log.Println("operation completed")
		return result
	}
}

//3rd party apis
func add(nos ...int) int {
	result := 0
	for i := 0; i < len(nos); i++ {
		result = nos[i]
	}
	return result
}

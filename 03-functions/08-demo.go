/* applied higher order functions */
package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	/*
		add(100,200)
		subtract(100,200)
	*/
	/*
		logAdd := getLogOperation(add)
		logAdd(100, 200)

		logSubtract := getLogOperation(subtract)
		logSubtract(100, 200)
	*/

	/*
		profileAdd := getProfileOperation(add)
		profileAdd(100, 200)

		profileSubtract := getProfileOperation(subtract)
		profileSubtract(100, 200)
	*/
	logAdd := getLogOperation(add)
	profileLogAdd := getProfileOperation(logAdd)
	profileLogAdd(100, 200)

	getProfileOperation(getLogOperation(subtract))(100, 200)
}

func getProfileOperation(operation func(int, int)) func(int, int) {
	return func(x, y int) {
		start := time.Now()
		operation(x, y)
		elapsed := time.Since(start)
		fmt.Println("Elapsed : ", elapsed)
	}
}

func getLogOperation(operation func(int, int)) func(int, int) {
	return func(x, y int) {
		log.Println("operation started")
		operation(x, y)
		log.Println("operation completed")
	}
}

//3rd party apis
func add(x, y int) {
	fmt.Println("Result :", x+y)
}

func subtract(x, y int) {
	fmt.Println("Result :", x-y)
}

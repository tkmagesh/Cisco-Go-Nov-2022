package main

import "fmt"

func main() {
	//convention - constants are name with all UPPERCASE
	//OK to have unused constants
	//const PI float32 = 3.14
	const PI = 3.14
	//PI = 2
	fmt.Println(PI)

	//iota
	/*
		const RED int = 0
		const GREEN int = 1
		const BLUE int = 2
	*/

	/*
		const (
			RED   int = 0
			GREEN int = 1
			BLUE  int = 2
		)
	*/

	/*
		const (
			RED   int = iota
			GREEN int = iota
			BLUE  int = iota
		)
	*/

	/*
		const (
			RED int = iota
			GREEN
			BLUE
		)
	*/

	/*
		const (
			RED = iota
			GREEN
			BLUE
		)
	*/

	/*
		const (
			RED   = iota + 2 //0 + 2
			GREEN            //1 + 2
			BLUE             //2 + 2
		)
	*/

	/*
		const (
			RED   = iota * 2 //0 * 2
			GREEN            //1 * 2
			BLUE             //2 * 2
		)
	*/

	const (
		RED   = iota * 2 //0 * 2
		GREEN            //1 * 2
		_                //2 * 2
		BLUE             //3 * 2
	)

	fmt.Printf("RED = %d, GREEN = %d, BLUE = %d\n", RED, GREEN, BLUE)

	/* iota applied */
	const (
		VERBOSE = 1 << iota
		CONFIG_FROM_DISK
		DATABASE_REQUIRED
		LOGGER_ACTIVATED
		DEBUG
		FLOAT_SUPPORT
		RECOVERY_MODE
		REBOOT_ON_FAILURE
	)
	fmt.Printf("%b, %b, %b, %b, %b, %b, %b, %b\n", VERBOSE, CONFIG_FROM_DISK, DATABASE_REQUIRED, LOGGER_ACTIVATED, DEBUG, FLOAT_SUPPORT, RECOVERY_MODE, REBOOT_ON_FAILURE)
}

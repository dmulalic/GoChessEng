package main

import "fmt"

func test() {
	fmt.Println("---   DEBUG   ---")
	for index := 0; index < BoardSquareNumber; index++ {
		if index%10 == 0 {
			fmt.Printf("\n")
		}
		fmt.Printf("%5d", Square120To64[index])
	}
	fmt.Println()
	fmt.Println()

	for index := 0; index < 64; index++ {
		if index%8 == 0 {
			fmt.Printf("\n")
		}
		fmt.Printf("%5d", Square64To120[index])
	}

	fmt.Printf("%v\n", NoSquare)
	fmt.Println("--- END DEBUG ---")
}

func main() {
	fmt.Printf("%v version %v\n", Name, Version)
	fmt.Println("MIT license")

	allInit()
	test()

}

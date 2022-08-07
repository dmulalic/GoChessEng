package main

import "fmt"

func main() {
	fmt.Printf("%v version %v\n", Name, Version)
	fmt.Println("MIT license")

	allInit()

	fmt.Println("---   DEBUG   ---")
	fmt.Printf("%v\n", NoSquare)
	fmt.Println("--- END DEBUG ---")
}

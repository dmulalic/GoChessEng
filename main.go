package main

import "fmt"

func main() {
	fmt.Printf("%v version %v\n", Name, Version)
	fmt.Println("MIT license")

	fmt.Println("\n ---   DEBUG   ---\n")
	fmt.Printf("%v", NoSquare)
	fmt.Println("\n --- END DEBUG ---\n")
}

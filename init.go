package main

import "fmt"

var (
	Square120To64 [BoardSquareNumber]int
	Square64To120 [64]int
)

func FileRank2Square(file int, rank int) int {

	return ((21 + file) + (rank * 10))
}

func allInit() {

	fmt.Println("allInit()")

}

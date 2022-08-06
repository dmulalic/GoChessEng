package main

import "fmt"

var (
	Name    string = "GoChessEng"
	Version string = "0.0.1"

// BoardSquareNumber int    = 120
)

// WP = white pawn, WN = white kinght, ...
const (
	Empty int = iota
	WP
	WN
	WB
	WR
	WQ
	WK
	BP
	BN
	BB
	BR
	BQ
	BK
)

func main() {
	fmt.Printf("%v version %v\n", Name, Version)
	fmt.Println("MIT license")
}

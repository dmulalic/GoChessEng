package main

import "fmt"

var (
	Name    string = "GoChessEng"
	Version string = "0.0.1"

// BoardSquareNumber int    = 120
)

const (
	Empty int = iota
	WhitePawn
	WhiteKnight
	WhiteBishop
	WhileRook
	WhiteQueen
	WhiteKing
	BlackPawn
	BlackKnight
	BlackBishop
	BlackRook
	BlackQueen
	BlackKing
)

func main() {
	fmt.Printf("%v version %v\n", Name, Version)
	fmt.Println("MIT license")
}

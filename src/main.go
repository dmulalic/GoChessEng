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

const (
	FileA int = iota
	FileB
	FileC
	FileD
	FileE
	FileF
	FileG
	FileH
	FileNone
)

const (
	Rank1 int = iota
	Rank2
	Rank3
	Rank4
	Rank5
	Rank6
	Rank7
	Rank8
	RankNone
)

func main() {
	fmt.Printf("%v version %v\n", Name, Version)
	fmt.Println("MIT license")
}

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

const (
	A1 int = iota + 21
	B1
	C1
	D1
	E1
	F1
	G1
	H1
)
const (
	A2 int = iota + 31
	B2
	C2
	D2
	E2
	F2
	G2
	H2
)
const (
	A3 int = iota + 41
	B3
	C3
	D3
	E3
	F3
	G3
	H3
)
const (
	A4 int = iota + 51
	B4
	C4
	D4
	E4
	F4
	G4
	H4
)
const (
	A5 int = iota + 61
	B5
	C5
	D5
	E5
	F5
	G5
	H5
)
const (
	A6 int = iota + 71
	B6
	C6
	D6
	E6
	F6
	G6
	H6
)
const (
	A7 int = iota + 81
	B7
	C7
	D7
	E7
	F7
	G7
	H7
)
const (
	A8 int = iota + 91
	B8
	C8
	D8
	E8
	F8
	G8
	H8
)

func main() {
	fmt.Printf("%v version %v\n", Name, Version)
	fmt.Println("MIT license")
}

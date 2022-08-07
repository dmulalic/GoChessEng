package main

var (
	Square120To64 [BoardSquareNumber]int
	Square64To120 [64]int
)

func FileRank2Square(file int, rank int) int {

	return ((21 + file) + (rank * 10))
}

func initSquare120to64() {
	var index int

	var file = FileA
	var rank = Rank1

	var square int = A1
	var square64 int

	for index = 0; index < BoardSquareNumber; index++ {
		Square120To64[index] = 65
	}

	for index = 0; index < 64; index++ {
		Square64To120[index] = 120
	}

	for rank = Rank1; rank <= Rank8; rank++ {
		for file = FileA; file <= FileH; file++ {
			square = FileRank2Square(file, rank)
			Square64To120[square64] = square
			Square120To64[square] = square64
			square64++
		}
	}
}

func allInit() {

	initSquare120to64()

}

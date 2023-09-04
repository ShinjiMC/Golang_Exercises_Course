package main

import "fmt"

// Declare a type named File which stores if a square is occupied by a piece - this will be a slice of bools
type File []bool

// Declare a type named Chessboard which contains a map of eight Files, accessed with keys from "A" to "H"
type Chessboard map[string]File

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file.
func CountInFile(cb Chessboard, file string) int {
	count := 0
	for _, occupied := range cb[file] {
		if occupied {
			count += 1
		}
	}
	return count
}

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank.
func CountInRank(cb Chessboard, rank int) int {
	if rank < 1 || rank > 8 {
		return 0
	}
	count := 0
	for _, file := range cb {
		if file[rank-1] {
			count += 1
		}
	}
	return count
}

// CountAll should count how many squares are present in the chessboard.
func CountAll(cb Chessboard) int {
	count := 0
	for _, file := range cb {
		for range file {
			count += 1
		}
	}
	return count
}

// CountOccupied returns how many squares are occupied in the chessboard.
func CountOccupied(cb Chessboard) int {
	count := 0
	for file := range cb {
		count += CountInFile(cb, file)
	}
	return count
}

func main() {
	board := Chessboard{
		"A": {true, false, false, false, true, false, false, true},
		"B": {false, false, false, false, false, false, false, false},
		"C": {false, false, false, false, true, false, false, true},
		"D": {false, true, false, false, false, false, false, true},
		"E": {false, false, false, false, false, false, true, true},
		"F": {true, false, true, false, false, false, false, true},
		"G": {false, false, false, false, false, false, false, true},
		"H": {true, false, false, false, false, false, false, true},
	}
	fmt.Println("Test 1: ", CountInFile(board, "A"))
	fmt.Println("Test 2: ", CountInRank(board, 2))
	fmt.Println("Test 3: ", CountAll(board))
	fmt.Println("Test 4: ", CountOccupied(board))
}

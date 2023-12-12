/* https://www.codewars.com/kata/525caa5c1bf619d28c000335

If we were to set up a Tic-Tac-Toe game, we would want to know whether the board's current state is solved, wouldn't we? Our goal is to create a function that will check that for us!

Assume that the board comes in the form of a 3x3 array, where the value is 0 if a spot is empty, 1 if it is an "X", or 2 if it is an "O", like so:

[[0, 0, 1],
 [0, 1, 2],
 [2, 1, 0]]

We want our function to return:

    -1 if the board is not yet finished AND no one has won yet (there are empty spots),
    1 if "X" won,
    2 if "O" won,
    0 if it's a cat's game (i.e. a draw).

You may assume that the board passed in is valid in the context of a game of Tic-Tac-Toe. */

package kata

const (
	pX = 1
	pO = 2
	p_ = 0

	WinX       = pX
	WinO       = pO
	Draw       = p_
	InProgress = -1
)

func checkWinning(line [3]int) int {
	const UnknownField = 99
	playerToken := UnknownField
	for _, elem := range line {
		if elem != p_ {
			if playerToken == UnknownField {
				playerToken = elem
			} else {
				if playerToken != elem {
					return InProgress
				}
			}
		} else {
			return InProgress
		}
	}
	return playerToken
}

func IsSolved(board [3][3]int) int {
	// check rows
	for _, row := range board {
		result := checkWinning(row)
		if result == WinO || result == WinX {
			return result
		}
	}
	// check columns
	var columns [3][3]int
	for idxRow, row := range board {
		for idxColumn, field := range row {
			columns[idxColumn][idxRow] = field
		}
	}

	for _, column := range columns {
		result := checkWinning(column)
		if result == WinO || result == WinX {
			return result
		}
	}
	// check diagonals
	diagonals := [2][3]int{
		{board[0][0], board[1][1], board[2][2]},
		{board[0][2], board[1][1], board[2][0]},
	}
	for _, diagonal := range diagonals {
		result := checkWinning(diagonal)
		if result == WinO || result == WinX {
			return result
		}
	}

	// check if the game is still ongoing
	for _, row := range board {
		for _, field := range row {
			if field == 0 {
				return InProgress
			}
		}
	}

	return Draw
}

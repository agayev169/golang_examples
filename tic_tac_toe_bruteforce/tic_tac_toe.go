package main

import "fmt"

const (
	ImpossibleMove = "impossible move"
)

const (
	EMPTY = iota
	X
	O
)

type TicTacToe struct {
	Board [][]int8
	Turn  int8
}

func NewTicTacToe() *TicTacToe {
	board := make([][]int8, 3)
	for i := 0; i < 3; i++ {
		board[i] = make([]int8, 3)
	}

	game := TicTacToe{
		Board: board,
		Turn:  X,
	}

	return &game
}

type TicTacToeMove struct {
	X int8
	Y int8
}

func (g TicTacToe) PossibleMoves() []Move {
	var res []Move
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if g.Board[i][j] == EMPTY {
				res = append(res, &TicTacToeMove{int8(j), int8(i)})
			}
		}
	}

	return res
}

func (g *TicTacToe) Move(m Move) error {
	move := m.(*TicTacToeMove)
	if g.Board[move.Y][move.X] == EMPTY {
		g.Board[move.Y][move.X] = g.Turn
		g.Turn = (g.Turn)%2 + 1 // X -> O, O -> x
		return nil
	}
	return fmt.Errorf(ImpossibleMove)
}

func (g *TicTacToe) GameOver() bool {
	if g.Board[0][0] != EMPTY && g.Board[0][0] == g.Board[1][1] && g.Board[1][1] == g.Board[2][2] {
		return true
	}
	if g.Board[0][2] != EMPTY && g.Board[0][2] == g.Board[1][1] && g.Board[1][1] == g.Board[2][0] {
		return true
	}
	for i := 0; i < 3; i++ {
		if g.Board[i][0] != EMPTY && g.Board[i][0] == g.Board[i][1] && g.Board[i][1] == g.Board[i][2] {
			return true
		}
		if g.Board[0][i] != EMPTY && g.Board[0][i] == g.Board[1][i] && g.Board[1][i] == g.Board[2][i] {
			return true
		}
	}
	if len(g.PossibleMoves()) == 0 {
		return true
	}
	return false
}

func (g *TicTacToe) Winner() int32 {
	if g.Board[0][0] == X && g.Board[0][0] == g.Board[1][1] && g.Board[1][1] == g.Board[2][2] {
		return X
	} else if g.Board[0][0] == O && g.Board[0][0] == g.Board[1][1] && g.Board[1][1] == g.Board[2][2] {
		return O
	}

	if g.Board[0][2] == X && g.Board[0][2] == g.Board[1][1] && g.Board[1][1] == g.Board[2][0] {
		return X
	} else if g.Board[0][2] == O && g.Board[0][2] == g.Board[1][1] && g.Board[1][1] == g.Board[2][0] {
		return O
	}

	for i := 0; i < 3; i++ {
		if g.Board[i][0] == X && g.Board[i][0] == g.Board[i][1] && g.Board[i][1] == g.Board[i][2] {
			return X
		} else if g.Board[i][0] == O && g.Board[i][0] == g.Board[i][1] && g.Board[i][1] == g.Board[i][2] {
			return O
		}

		if g.Board[0][i] == X && g.Board[0][i] == g.Board[1][i] && g.Board[1][i] == g.Board[2][i] {
			return X
		} else if g.Board[0][i] == O && g.Board[0][i] == g.Board[1][i] && g.Board[1][i] == g.Board[2][i] {
			return O
		}
	}

	return EMPTY
}

func (g *TicTacToe) String() string {
	res := ""
	for i := 0; i < 3; i++ {
		res += "|"
		for j := 0; j < 3; j++ {
			if g.Board[j][i] == X {
				res += "X"
			} else if g.Board[j][i] == O {
				res += "O"
			} else {
				res += " "
			}
			res += "|"
		}

		if i != 2 {
			res += "\n"
		}
	}
	return res
}

func CopyTicTacToe(g *TicTacToe) *TicTacToe {
	newGame := NewTicTacToe()
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			newGame.Board[i][j] = g.Board[i][j]
		}
	}
	newGame.Turn = g.Turn

	return newGame
}

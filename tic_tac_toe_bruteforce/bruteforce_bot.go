package main

type TicTacToeBot struct {
	Side int8
}

func (b TicTacToeBot) Move(g Game) Move {
	m, _ := b.miniMax(g)
	return m
}

func (b TicTacToeBot) miniMax(g Game) (*TicTacToeMove, int8) {
	gameOver := g.GameOver()
	if gameOver {
		winner := g.Winner()
		if int8(winner) == b.Side {
			return nil, 1
		} else if int8(winner%2+1) == b.Side {
			return nil, -1
		} else {
			return nil, 0
		}
	}

	var bestMove *TicTacToeMove = nil
	var bestScore int8 = 0
	game := g.(*TicTacToe)
	pm := g.PossibleMoves()
	for _, m := range pm {
		move := m.(*TicTacToeMove)
		err := game.Move(m)
		if err != nil {
			panic(err)
		}

		_, score := b.miniMax(game)
		if score > bestScore || bestMove == nil {
			bestScore = score
			bestMove = move
		}
		game.Board[move.Y][move.X] = EMPTY
	}

	return bestMove, bestScore
}

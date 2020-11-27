package main

import "fmt"

func main() {
	totalGames := 100
	wins := 0
	loses := 0

	bots := make([]Bot, 2)
	botIdx := 0

	bots[0] = &TicTacToeBot{X}
	bots[1] = &RandomBot{}

	for i := 0; i < totalGames; i++ {
		g := NewTicTacToe()
		for !g.GameOver() {
			m := bots[botIdx].Move(g)
			err := g.Move(m)
			if err != nil {
				panic(err)
			}

			botIdx = (botIdx + 1) % 2
		}
		winner := g.Winner()
		if winner == X {
			wins++
		} else if winner == O {
			loses++
		}
		botIdx = 0
	}

	fmt.Printf("Wins: %v, loses: %v\n", wins, loses)
}
